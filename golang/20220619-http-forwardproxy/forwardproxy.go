package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

type Proxy struct {
}

func (p *Proxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Printf("recv req: %v %v from %v", req.Method, req.Host, req.RemoteAddr)
	transport := http.DefaultTransport

	// 复制请求
	outReq := new(http.Request)
	*outReq = *req

	// 获取IP
	clientIP, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		log.Printf("net.SplitHostPort(%v) error: %v", req.RemoteAddr, err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	prior, ok := outReq.Header["X-Forwarded-For"]
	if ok {
		clientIP = strings.Join(prior, ", ") + ", " + clientIP
	}
	log.Printf("clientip: %v", clientIP)
	outReq.Header.Set("X-Forwarded-For", clientIP)

	// 请求
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}

	// 返回响应
	for k, v := range res.Header {
		for _, vv := range v {
			rw.Header().Add(k, vv)
		}
	}
	rw.WriteHeader(res.StatusCode)
	io.Copy(rw, res.Body)
	res.Body.Close()
}

func main() {
	http.Handle("/", &Proxy{})
	http.ListenAndServe(":8080", nil)
}
