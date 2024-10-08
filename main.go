package main

import (
	"crypto/tls"
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"os"
	"sync"
	"strconv"
	"strings"
	"time"
	"github.com/google/uuid"
)

var (
	client = &http.Client{}
	version = "unknown"
	commit = "unknown"
	date = "unknown"
)

var opts Options

func httpGet(reqUrl url.URL, thread, counter, total int) {
	if opts.UUID {
		values := reqUrl.Query()
		u, err := uuid.NewRandom()
		if err != nil {
			log.Fatalf("Failed to generate UUID: %s", err)
		}
		values.Set("uuid", u.String())
		reqUrl.RawQuery = values.Encode()
	}
	urlStr := reqUrl.String()
	request, err := http.NewRequest("GET", urlStr, nil)
	if opts.Trace > 0 {
		trace := &httptrace.ClientTrace{
			DNSStart: traceDNSStart,
			DNSDone: traceDNSDone,
			GetConn: traceGetConn,
			GotConn: traceGotConn,
			ConnectStart: traceConnectStart,
			ConnectDone: traceConnectDone,
			TLSHandshakeDone: traceTLSHandshakeDone,
			PutIdleConn: tracePutIdleConn,
		}
		request = request.WithContext(httptrace.WithClientTrace(request.Context(), trace))
	}
	if err != nil {
		log.Printf("[%03d-%05d] ERROR at http.NewRequest: %v %s\n", thread, counter, err, urlStr)
		return
	}
	start := time.Now()
	resp, err := client.Do(request)
	if err != nil {
		log.Printf("[%03d-%05d] ERROR %d ms at Do(request): %v %s\n", thread, counter, time.Since(start).Milliseconds(), err, urlStr)
		return
	}
	defer func() {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()
	reads, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		log.Printf("[%03d-%05d] ERROR read %d bytes, %s %s\n", thread, counter, reads, err, urlStr)
	}
	resp.Body.Close()
	end := time.Now()
	if opts.Trace >= 3 {
		log.Printf("[%03d-%05d] DEBUG response header: %#v, %s\n", thread, counter, resp.Header, urlStr)
	}
	diff := end.Sub(start).Milliseconds()
	if diff > int64(opts.ShowThresholdMs) {
		log.Printf("[%03d-%05d] WARN %03d %5d ms, %d bytes, %s\n", thread, counter, resp.StatusCode, diff, reads, urlStr)
	} else if opts.Verbose {
		log.Printf("[%03d-%05d] INFO %03d %5d ms, %d bytes, %s\n", thread, counter, resp.StatusCode, diff, reads, urlStr)
	}
	if ! opts.Verbose && counter > 0 && counter % 100 == 0 {
		log.Printf("[%03d-%05d] INFO %d/%d requests finished\n", thread, counter, counter, total)
	}
}

func httpGetWithRandomSleep(reqUrl url.URL, thread, count, sleepMin, sleepMax int, wg *sync.WaitGroup) {
	for i := 1; i <= count; i++ {
		if i > 1 && (sleepMin > 0 || sleepMax > 0) {
			if sleepMax == sleepMin {
				time.Sleep(time.Duration(sleepMin) * time.Millisecond)
			} else {
				time.Sleep(time.Duration(sleepMin + rand.Intn(sleepMax - sleepMin)) * time.Millisecond)
			}
		}
		httpGet(reqUrl, thread, i, count)
	}
	if opts.SleepAtEnd {
		if (sleepMin > 0 || sleepMax > 0) {
			if sleepMax == sleepMin {
				time.Sleep(time.Duration(sleepMin) * time.Millisecond)
			} else {
				time.Sleep(time.Duration(sleepMin + rand.Intn(sleepMax - sleepMin)) * time.Millisecond)
			}
		}
	}
	wg.Done()
}

type Options struct {
	Requests               int  `short:"c" long:"requests" default:"1" description:"Number of requests per thread."`
	Threads                int  `short:"t" long:"threads" default:"1" description:"Number of threads."`
	ConnectTimeoutSec      int  `long:"connect-timeout" default:"10" description:"Connect timeout in second."`
	TLSHandshakeTimeoutSec int  `long:"tls-handshake-timeout" default:"10" description:"TLS handshake timeout in second."`
	MaxIdleConns           int  `long:"max-idle-conns" default:"2" description:"Max idle connections. Zero means no limit. Override with max-idle-conns-per-host if max-idle-conns-per-host is greater than max-idle-conns"`
	MaxIdleConnsPerHost    int  `long:"max-idle-conns-per-host" default:"2" description:"Max idle connections per host."`
	MaxConnsPerHost        int  `long:"max-conns-per-host" default:"10" description:"Max connections per host. Zero means no limit."`
	IdleConnTimeoutSec     int  `long:"idle-conn-timeout" default:"60" description:"Idle connection timeout in second."`
	KeepAliveIntervalSec   int  `long:"tcp-keepalive-interval" default:"0" description:"TCP keepalive interval in second. Zero means 15 seconds."`
	DisableKeepAlives      bool `long:"disable-http-keepalive" description:"Disable HTTP Keep-Alive. "`
	Insecure               bool `short:"k" long:"insecure" description:"Skip TLS cert verify."`
	TimeoutSec             int  `short:"T" long:"timeout" default:"30" description:"Request total timeout in second."`
	Version                bool `short:"V" long:"version" description:"Show version and exit."`
	Verbose                bool `short:"v" long:"verbose" description:"Enable verbose output. Show response time every request."`
	ShowThresholdMs        int  `short:"s" long:"show-threshold" default:"200" description:"Show response time in Millisecond if over this threshold."`
	SleepMaxMs             int  `short:"r" long:"random-sleep-max-ms" default:"0" description:"Max interval sleep time in millisecond. (DEPRECATED)"`
	SleepRangeMs           string `short:"S" long:"sleep-range-ms" default:"0:0" description:"Range of andom sleep time (min:max) in millisecond."`
	SleepAtEnd             bool `long:"sleep-at-end" description:"Sleep at end."`
	ServerName             string `long:"servername" description:"Server Name Indication extension in TLS handshake."`
	Trace                  int  `long:"trace" description:"Set httptrace log level in (1,2,3). The Larger, more verbose."`
	UUID                   bool `long:"uuid" description:"Add uuid query string to identify each request."`
	Args                   struct {
		Url string `description:"URL"`
	} `positional-args:"yes"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	if opts.Version {
		fmt.Printf("version: %s\ncommit: %s\ndate: %s\n", version, commit, date)
		os.Exit(0)
	}

	if opts.Args.Url == "" {
		log.Fatal("url parameter required")
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: opts.Insecure,
	}

	if opts.ServerName != "" {
		tlsConfig.ServerName = opts.ServerName
	}

	if opts.MaxIdleConns > 0 && opts.MaxIdleConnsPerHost > opts.MaxIdleConns {
		opts.MaxIdleConns = opts.MaxIdleConnsPerHost
	}

	sleepMin := 0
	sleepMax := 0

	if strings.Contains(opts.SleepRangeMs, ":") {
		s := strings.Split(opts.SleepRangeMs, ":")
		sleepMin, err = strconv.Atoi(s[0])
		if err != nil {
			log.Fatalf("Invalid sleep range: %v", opts.SleepRangeMs)
		}
		sleepMax, err = strconv.Atoi(s[1])
		if err != nil {
			log.Fatalf("Invalid sleep range: %v", opts.SleepRangeMs)
		}
	}
	if opts.SleepMaxMs > 0 {
		sleepMin = 0
		sleepMax = opts.SleepMaxMs
	}

	reqUrl, err := url.ParseRequestURI(opts.Args.Url)
	if err != nil {
		log.Fatalf("Failed to ParseRequestURL(%s): %s", opts.Args.Url, err)
	}

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   time.Duration(opts.ConnectTimeoutSec) * time.Second,
			KeepAlive: time.Duration(opts.KeepAliveIntervalSec) * time.Second,
			DualStack: true,
		}).DialContext,
		IdleConnTimeout:     time.Duration(opts.IdleConnTimeoutSec) * time.Second,
		MaxIdleConns:        opts.MaxIdleConns,
		MaxIdleConnsPerHost: opts.MaxIdleConnsPerHost,
		MaxConnsPerHost:     opts.MaxConnsPerHost,
		TLSHandshakeTimeout: time.Duration(opts.TLSHandshakeTimeoutSec) * time.Second,
		TLSClientConfig: tlsConfig,
		DisableKeepAlives: opts.DisableKeepAlives,
	}

	if os.Getenv("HTTP_PROXY") != "" {
		proxyUrl, err := url.Parse(os.Getenv("HTTP_PROXY"))
		if err != nil {
			log.Printf("WARN invalid HTTP_PROXY: %v", err)
		} else {
			transport.Proxy = http.ProxyURL(proxyUrl)
		}
	}

	client.Transport = transport

	client.Timeout = time.Duration(opts.TimeoutSec) * time.Second

	wg := sync.WaitGroup{}
	wg.Add(opts.Threads)

	for i := 0; i < opts.Threads; i++ {
		go httpGetWithRandomSleep(*reqUrl, i, opts.Requests, sleepMin, sleepMax, &wg)
	}

	wg.Wait()

	client.CloseIdleConnections()
}
