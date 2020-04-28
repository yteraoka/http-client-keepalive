# A tool to reproduce network trouble in some specific case.

```
Usage:
  http-client-keepalive [OPTIONS] [Url]

Application Options:
  -c, --requests=                Number of requests per thread. (default: 1)
  -t, --threads=                 Number of threads. (default: 1)
      --connect-timeout=         Connect timeout in second. (default: 10)
      --tls-handshake-timeout=   TLS handshake timeout in second. (default: 10)
      --max-idle-conns=          Max idle connections. Zero means no limit. (default: 2)
      --max-idle-conns-per-host= Max idle connections per host. (default: 2)
      --max-conns-per-host=      Max connections per host. Zero means no limit. (default: 10)
      --idle-conn-timeout=       Idle connection timeout in second. (default: 60)
      --tcp-keepalive-interval=  TCP keepalive interval in second. Zero means 15 seconds. (default: 0)
      --disable-http-keepalive   Disable HTTP Keep-Alive.
  -k, --insecre                  Skip TLS cert verify.
  -T, --timeout=                 Request total timeout in second. (default: 30)
  -V, --version                  Show version and exit.
  -v, --verbose                  Enable verbose output. Show response time every request.
  -s, --show-threshold=          Show response time in Millisecond if over this threshold. (default: 200)
  -r, --random-sleep-max-ms=     Max interval sleep time in millisecond. (default: 0)
      --servername=              Server Name Indication extension in TLS handshake.

Help Options:
  -h, --help                     Show this help message

Arguments:
  Url:                           URL
```
