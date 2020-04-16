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

func httpGet(url string, thread, counter int) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[%03d-%05d] ERROR in http.NewRequest: %v\n", thread, counter, err)
		return
	}
	start := time.Now()
	resp, err := client.Do(request)
	if err != nil {
		log.Printf("[%03d-%05d] ERROR in Do(request): %v\n", thread, counter, err)
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
	if opts.Verbose || diff > int64(opts.ShowThresholdMs) {
		log.Printf("[%03d-%05d] %d %d ms\n", thread, counter, resp.StatusCode, diff)
	}
}

func httpGetWithRandomSleep(url string, thread, count int, wg *sync.WaitGroup) {
	for i := 0; i < count; i++ {
		time.Sleep(time.Duration(rand.Intn(opts.SleepMaxMs)) * time.Millisecond)
		httpGet(url, thread, i)
	}
	wg.Done()
}

type Options struct {
	Requests               int  `short:"c" long:"requests" default:"1" description:"Number of requests per thread."`
	Threads                int  `short:"t" long:"threads" default:"1" description:"Number of threads."`
	ConnectTimeoutSec      int  `long:"connect-timeout" default:"10" description:"Connect timeout in second."`
	TLSHandshakeTimeoutSec int  `long:"tls-handshake-timeout" default:"10" description:"TLS handshake timeout in second."`
	MaxIdleConns           int  `long:"max-idle-conns" default:"2" description:"Max idle connections. Zero means no limit."`
	MaxIdleConnsPerHost    int  `long:"max-idle-conns-per-host" default:"2" description:"Max idle connections per host."`
	MaxConnsPerHost        int  `long:"max-conns-per-host" default:"10" description:"Max connections per host. Zero means no limit."`
	IdleConnTimeoutSec     int  `long:"idle-conn-timeout" default:"60" description:"Idle connection timeout in second."`
	KeepAliveIntervalSec   int  `long:"tcp-keepalive-interval" default:"0" description:"TCP keepalive interval in second."`
	Insecure               bool `short:"k" long:"insecre" description:"Skip TLS cert verify."`
	TimeoutSec             int  `short:"T" long:"timeout" default:"30" description:"Request total timeout in second."`
	Version                bool `short:"V" long:"version" description:"Show version and exit."`
	Verbose                bool `short:"v" long:"verbose" description:"Enable verbose output. Show response time every request."`
	ShowThresholdMs        int  `short:"s" long:"show-threshold" default:"200" description:"Show response time in Millisecond if over this threshold."`
	SleepMaxMs             int  `short:"r" long:"random-sleep-max-ms" default:"1000" description:"Max interval sleep time in millisecond."`
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
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: opts.Insecure,
		},
	}
	client.Timeout = time.Duration(opts.TimeoutSec) * time.Second

	wg := sync.WaitGroup{}
	wg.Add(opts.Threads)

	for i := 0; i < opts.Threads; i++ {
		go httpGetWithRandomSleep(opts.Args.Url, i, opts.Requests, &wg)
	}

	wg.Wait()
}
