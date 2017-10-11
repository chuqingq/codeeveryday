$ ./wrk -c 100 -t 10 -d 10 http://127.0.0.1:8081/hello
\Running 10s test @ http://127.0.0.1:8081/hello
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   197.95us  346.16us  19.81ms   97.55%
    Req/Sec    53.35k     4.47k   89.28k    95.53%
  5340664 requests in 10.10s, 677.40MB read
Requests/sec: 528795.89
Transfer/sec:     67.07MB

