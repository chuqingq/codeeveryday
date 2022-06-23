## 目标

抓取https明文。例如访问https://www.baidu.com/，获取明文数据。

## 架构

browser -> proxy -> fake_https_server -> https_server(https://www.baidu.com/)

## 方案

1. 用上节的CA证书签名CN为www.baidu.com的http证书：http.crt。
2. proxy接收到CONNECT后：用http.key、http.crt启动tls server；启动tls client连上www.baidu.com:443。然后双向拷贝（会自动handshake）
3. 如果要抓取browser发送给www.baidu.com的内容，可以对tls server的write进行封装，打印出内容。例如：

```go
2022/06/23 13:57:23 write: GET /s?ie=utf-8&f=8&rsv_bp=1&rsv_idx=2&ch=&tn=baiduhome_pg&bar=&wd=chuqq&rsv_spt=1&oq=123&rsv_pq=bd78bb670012f90e&rsv_t=8422hiK2GgW1OU%2FjUNSaHNJRKYjGIM24zH213VYFreF8vpyI7woTDO0k4UmRNi1flJPr&rqlang=cn HTTP/1.1
Host: www.baidu.com
...
```

其中wd=chuqq是百度搜索的关键字。
