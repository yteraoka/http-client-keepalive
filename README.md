# A tool to reproduce network trouble in some specific case.

```
Usage:
  http-client-keepalive [OPTIONS] [Url]

Application Options:
  -c, --requests=                Number of requests per thread. (default: 1)
  -t, --threads=                 Number of threads. (default: 1)
      --connect-timeout=         Connect timeout in second. (default: 10)
      --tls-handshake-timeout=   TLS handshake timeout in second. (default: 10)
      --max-idle-conns=          Max idle connections. Zero means no limit. Override with max-idle-conns-per-host if
                                 max-idle-conns-per-host is greater than max-idle-conns (default: 2)
      --max-idle-conns-per-host= Max idle connections per host. (default: 2)
      --max-conns-per-host=      Max connections per host. Zero means no limit. (default: 10)
      --idle-conn-timeout=       Idle connection timeout in second. (default: 60)
      --tcp-keepalive-interval=  TCP keepalive interval in second. Zero means 15 seconds. (default: 0)
      --disable-http-keepalive   Disable HTTP Keep-Alive.
  -k, --insecure                 Skip TLS cert verify.
  -T, --timeout=                 Request total timeout in second. (default: 30)
  -V, --version                  Show version and exit.
  -v, --verbose                  Enable verbose output. Show response time every request.
  -s, --show-threshold=          Show response time in Millisecond if over this threshold. (default: 200)
  -r, --random-sleep-max-ms=     Max interval sleep time in millisecond. (DEPRECATED) (default: 0)
  -S, --sleep-range-ms=          Range of andom sleep time (min:max) in millisecond. (default: 0:0)
      --servername=              Server Name Indication extension in TLS handshake.
      --trace=                   Set httptrace log level in (1,2,3). The Larger, more verbose.

Help Options:
  -h, --help                     Show this help message

Arguments:
  Url:                           URL
```

use HTTP Proxy

```
HTTP_PROXY=http://192.168.1.1:3128 http-client-keepalive https://example.com/
HTTP_PROXY=http://user:password@192.168.1.1:3128 http-client-keepalive https://example.com/
```
