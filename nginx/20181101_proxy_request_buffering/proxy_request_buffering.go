package main

import (
    "net/http"
	"fmt"
	"log"
)

var response = []byte("hello world")

func HandleFunc(res http.ResponseWriter, req * http.Request) {
    log.Printf("HandleFunc()")
    res.Write(response)
}

func main() {
    http.HandleFunc("/", HandleFunc)
    err := http.ListenAndServe(":8081", nil)
    if err != nil {
        fmt.Println(err)
    }
}


/*
nginx配置：
        location /proxy_request_buffering_off {
                proxy_http_version 1.1;
                proxy_request_buffering off;
                proxy_pass http://127.0.0.1:8081/;
        }

        location /proxy_request_buffering_on {
                proxy_http_version 1.1;
                proxy_request_buffering on;
                proxy_pass http://127.0.0.1:8081/;
        }

验证：

场景1：开启proxy_request_buffering的场景。结果：敲完67890后的回车才结束。

GET /proxy_request_buffering_on HTTP/1.1
Host: 127.0.0.1
Content-Length: 12

12345
67890

场景2：关闭proxy_request_buffering。结果：在敲Content-Length后第二个回车后就结束了。因为都到HandleFunc，直接返回响应了
GET /proxy_request_buffering_off HTTP/1.1
Host: 127.0.0.1
Content-Length: 12

12345
67890
*/