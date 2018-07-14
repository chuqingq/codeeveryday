package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

var response = []byte("hello world")

func main() {
	err := fasthttp.ListenAndServe(":8081", func(ctx *fasthttp.RequestCtx) {
		ctx.SetBody(response)
	})
	if err != nil {
		log.Printf("error: %v", err)
	}
}
