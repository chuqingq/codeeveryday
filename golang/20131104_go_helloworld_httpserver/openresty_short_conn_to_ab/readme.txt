
`echo-nginx-module`[](https://github.com/openresty/echo-nginx-module) is used.

# wrk can enable `Connection: keep-alive`

if i use wrk `time ./wrk -c 100 -d 20 -t 4 http://127.0.0.1:8081/`, the captured stream is as follows:

```
GET / HTTP/1.1
Host: 127.0.0.1:8081

HTTP/1.1 200 OK
Server: openresty/1.11.2.5
Date: Tue, 19 Sep 2017 00:44:05 GMT
Content-Type: application/octet-stream
Transfer-Encoding: chunked
Connection: keep-alive

c
hello world

0

GET / HTTP/1.1
Host: 127.0.0.1:8081

HTTP/1.1 200 OK
Server: openresty/1.11.2.5
Date: Tue, 19 Sep 2017 00:44:05 GMT
Content-Type: application/octet-stream
Transfer-Encoding: chunked
Connection: keep-alive

c
hello world

0
```

the result of `netstat -anp|grep 8081|wc -l` is 208, include 8 listening sockets and 100*2 establised sockets.


# ab cannot enable `Connection: keep-alive`

if i use ab without `-k`, ab will not reuse the connection;
if i use ab whth `-k`, ab will get response stream as follows:

```
GET / HTTP/1.0
Connection: Keep-Alive
Host: 127.0.0.1:8081
User-Agent: ApacheBench/2.3
Accept: */*

HTTP/1.1 200 OK
Server: openresty/1.11.2.5
Date: Tue, 19 Sep 2017 00:42:13 GMT
Content-Type: application/octet-stream
Connection: close

hello world

```

the result of `netstat -anp|grep 8081|wc -l` will exceed to more than 2000 sockets.
I think the reason is, ab use `HTTP/1.0`.

