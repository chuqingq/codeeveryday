require "socket"

s = TCPServer.new '127.0.0.1', 8081
s.listen 10000

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

=begin
$ time ./wrk -c 100 -d 20 -t 4 http://127.0.0.1:8081/
Running 20s test @ http://127.0.0.1:8081/
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.29ms  185.61us   5.28ms   87.20%
    Req/Sec    19.17k     1.64k   20.94k    83.00%
  1525976 requests in 20.00s, 139.71MB read
Requests/sec:  76290.25
Transfer/sec:      6.98MB

real	0m20.019s
user	0m4.380s
sys	0m39.856s
=end

