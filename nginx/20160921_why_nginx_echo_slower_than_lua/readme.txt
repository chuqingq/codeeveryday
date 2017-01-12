/hellolua是lua，/hello是echo。

结论：echo模块区分http版本，如果是1.1，则可以keepalive，如果是1.0则不论是否带keepalive头域，都是短连接。

root@Offline-Basic-002:/home/imax/xpush/xpush_1.9.7/access/lua# curl http://127.0.0.1:58080/hellolua -v -0 -H "Connection: keep-alive"
* Hostname was NOT found in DNS cache
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 58080 (#0)
> GET /hellolua HTTP/1.0
> User-Agent: curl/7.35.0
> Host: 127.0.0.1:58080
> Accept: */*
> Connection: keep-alive
> 
< HTTP/1.1 200 OK
* Server openresty/1.9.15.1 is not blacklisted
< Server: openresty/1.9.15.1
< Date: Wed, 21 Sep 2016 04:18:55 GMT
< Content-Type: application/octet-stream
< Content-Length: 12
< Connection: keep-alive
< 
hello world
* Connection #0 to host 127.0.0.1 left intact
root@Offline-Basic-002:/home/imax/xpush/xpush_1.9.7/access/lua# curl http://127.0.0.1:58080/hello -v -0 -H "Connection: keep-alive"
* Hostname was NOT found in DNS cache
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 58080 (#0)
> GET /hello HTTP/1.0
> User-Agent: curl/7.35.0
> Host: 127.0.0.1:58080
> Accept: */*
> Connection: keep-alive
> 
< HTTP/1.1 200 OK
* Server openresty/1.9.15.1 is not blacklisted
< Server: openresty/1.9.15.1
< Date: Wed, 21 Sep 2016 04:19:05 GMT
< Content-Type: application/octet-stream
< Connection: close
< 
hello world
* Closing connection 0
