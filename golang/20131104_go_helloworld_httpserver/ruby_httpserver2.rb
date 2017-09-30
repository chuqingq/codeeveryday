require "socket"

s = TCPServer.new '127.0.0.1', 8081
s.listen 10000

6.times do
  fork {
    loop do
      io = s.accept
      content = 'hello world'
      io << "HTTP/1.1 200 OK
Content-Type: text/html; utf-8
Connection: close
Content-Length: #{content.size}

#{content}"
      io.close
    end
  }
end
Process.waitall

=begin
$ time ./wrk -c 100 -d 20 -t 4 http://127.0.0.1:8081/
Running 20s test @ http://127.0.0.1:8081/
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   197.60us   85.20us   8.68ms   77.10%
    Req/Sec    31.57k     1.59k   45.45k    92.26%
  2515481 requests in 20.10s, 230.30MB read
Requests/sec: 125149.10
Transfer/sec:     11.46MB

real	0m20.103s
user	0m6.704s
sys	1m12.608s
=end

