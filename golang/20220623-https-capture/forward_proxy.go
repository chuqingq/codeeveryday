package main

import (
	"crypto/tls"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

type writerWrapper struct {
	io.Writer
}

func (w *writerWrapper) Write(p []byte) (n int, err error) {
	log.Printf("write: %v", string(p))
	n, err = w.Writer.Write(p)
	return
}

func serveConnect(w http.ResponseWriter, r *http.Request) {
	// 回复成功
	w.WriteHeader(http.StatusOK)

	// hijack
	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}
	conn, _, err := hj.Hijack() // _bufrw
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Don't forget to close the connection:
	defer conn.Close()

	// 启动TLS server，并和客户端握手完成
	cert, err := tls.LoadX509KeyPair("./http.crt", "./http.key")
	if err != nil {
		log.Printf("tls loadkeypair error: %v", err)
		return
	}
	serverTLSConfig := tls.Config{Certificates: []tls.Certificate{cert}}
	serverTLSConn := tls.Server(conn, &serverTLSConfig)

	// 启动TLS client，并和server握手完成
	clientConn, err := tls.Dial("tcp", "www.baidu.com:443", &tls.Config{})
	if err != nil {
		log.Printf("tilsDial() error: %v", err)
		return
	}
	// io.copy 只看发送的内容，即client到server的内容。
	go io.Copy(&writerWrapper{Writer: clientConn}, serverTLSConn)
	io.Copy(serverTLSConn, clientConn)
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	log.Printf("recv req: %v %v from %v", r.Method, r.Host, r.RemoteAddr)
	// https CONNECT
	if r.Method == http.MethodConnect {
		serveConnect(w, r)
		return
	}
	// http
	transport := http.DefaultTransport

	// 复制请求
	outReq := new(http.Request)
	*outReq = *r

	// 获取IP
	clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Printf("ERROR net.SplitHostPort(%v) error: %v", r.RemoteAddr, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	prior, ok := outReq.Header["X-Forwarded-For"]
	if ok {
		clientIP = strings.Join(prior, ", ") + ", " + clientIP
	}
	// log.Printf("clientip: %v", clientIP)
	outReq.Header.Set("X-Forwarded-For", clientIP)

	// 请求
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		log.Printf("ERROR StatusBadGateway error: %v", err)
		return
	}

	// 返回响应
	for k, v := range res.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
	res.Body.Close()
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.HandleFunc("www.baidu.com:443", handleFunc)
	addr := ":8080"
	log.Printf("serve at: %v", addr)
	http.ListenAndServe(addr, nil)
}
