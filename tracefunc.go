package main

import (
	"crypto/tls"
	"log"
	"net/http/httptrace"
	"strings"
)

func traceDNSStart(info httptrace.DNSStartInfo) {
	if opts.Trace >= 3 {
		log.Printf("DNSStart: %+v\n", info.Host)
	}
}

func traceDNSDone(info httptrace.DNSDoneInfo) {
	if info.Err != nil {
		if opts.Trace >= 1 {
			log.Printf("DNSDone: %+v\n", info.Err)
		}
	} else {
		if opts.Trace >= 3 {
			log.Printf("DNSDone: addrs:%+v, Coalesced:%v\n", info.Addrs, info.Coalesced)
		}
	}
}

func traceGetConn(hostPort string) {
	if opts.Trace >= 3 {
		log.Printf("GetConn: %+v\n", hostPort)
	}
}

func traceGotConn(connInfo httptrace.GotConnInfo) {
	if opts.Trace >= 2 {
		log.Printf("GotConn: LocalAddr:%v, RemoteAddr:%v, Reused:%v, WasIdle:%v, IdlwTime:%dms",
				connInfo.Conn.LocalAddr(),
				connInfo.Conn.RemoteAddr(),
				connInfo.Reused,
				connInfo.WasIdle,
				connInfo.IdleTime.Milliseconds())
	}
}

func traceConnectStart(network, addr string) {
	if opts.Trace >= 3 {
		log.Printf("ConnectStart: connecting to %v %v\n", network, addr)
	}
}

func traceConnectDone(network, addr string, err error) {
	if err == nil {
		if opts.Trace >= 3 {
			log.Printf("ConnectDone: connected to %v %v\n", network, addr)
		}
	} else {
		log.Printf("ConnectDone: connected to %v %v, error:%v\n", network, addr, err)
	}
}

func traceTLSHandshakeDone(state tls.ConnectionState, err error) {
	if err == nil {
		if opts.Trace >= 3 {
			log.Printf("TLSHandshakeDone: version:%v, complete:%v, resume:%v\n",
					state.Version,
					state.HandshakeComplete,
					state.DidResume)
		}
	} else {
		log.Printf("TLSHandshakeDone: error:%v\n", err)
	}
}

func tracePutIdleConn(err error) {
	if err == nil {
		return
	}
	if strings.Contains(err.Error(), "too many idle connections for host") {
		if opts.Trace >= 2 {
			log.Printf("PutIdleConn: error:%v\n", err)
		}
	} else {
		log.Printf("PutIdleConn: error:%v\n", err)
	}
}
