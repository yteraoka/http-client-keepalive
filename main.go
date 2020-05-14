package main

import (
	"crypto/tls"
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	client = &http.Client{}
	version = "unknown"
	commit = "unknown"
	date = "unknown"
)

var opts Options

func httpGet(url string, thread, counter, total int) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[%03d-%05d] ERROR at http.NewRequest: %v\n", thread, counter, err)
		return
	}
	start := time.Now()
	resp, err := client.Do(request)
	if err != nil {
		log.Printf("[%03d-%05d] ERROR %d ms at Do(request): %v\n", thread, counter, time.Now().Sub(start).Milliseconds(), err)
		return
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	end := time.Now()
	diff := end.Sub(start).Milliseconds()
	if diff > int64(opts.ShowThresholdMs) {
		log.Printf("[%03d-%05d] WARN %d %d ms\n", thread, counter, resp.StatusCode, diff)
	} else if opts.Verbose {
		log.Printf("[%03d-%05d] INFO %d %d ms\n", thread, counter, resp.StatusCode, diff)
	}
	if ! opts.Verbose && counter > 0 && counter % 100 == 0 {
		log.Printf("[%03d-%05d] INFO %d/%d requests finished\n", thread, counter, counter, total)
	}
}

func httpGetWithRandomSleep(url string, thread, count int, wg *sync.WaitGroup) {
	for i := 1; i <= count; i++ {
		if opts.SleepMaxMs > 0 {
			time.Sleep(time.Duration(rand.Intn(opts.SleepMaxMs)) * time.Millisecond)
		}
		httpGet(url, thread, i, count)
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
	Insecure               bool `short:"k" long:"insecre" description:"Skip TLS cert verify."`
	TimeoutSec             int  `short:"T" long:"timeout" default:"30" description:"Request total timeout in second."`
	Version                bool `short:"V" long:"version" description:"Show version and exit."`
	Verbose                bool `short:"v" long:"verbose" description:"Enable verbose output. Show response time every request."`
	ShowThresholdMs        int  `short:"s" long:"show-threshold" default:"200" description:"Show response time in Millisecond if over this threshold."`
	SleepMaxMs             int  `short:"r" long:"random-sleep-max-ms" default:"0" description:"Max interval sleep time in millisecond."`
	ServerName             string `long:"servername" description:"Server Name Indication extension in TLS handshake."`
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

	client.Transport = &http.Transport{
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
	client.Timeout = time.Duration(opts.TimeoutSec) * time.Second

	wg := sync.WaitGroup{}
	wg.Add(opts.Threads)

	for i := 0; i < opts.Threads; i++ {
		go httpGetWithRandomSleep(opts.Args.Url, i, opts.Requests, &wg)
	}

	wg.Wait()
}
