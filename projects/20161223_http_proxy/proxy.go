package main

import (
	"flags"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

func main() {
	listen := flag.String("listen", ":8080", "listen address")
	flag.Parse()

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
