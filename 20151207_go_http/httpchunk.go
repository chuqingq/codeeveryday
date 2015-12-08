package main

import (
	"fmt"
	"net/http"
)

func HandleFunc(res http.ResponseWriter, req *http.Request) {
	// res.Header().Set("Content-Length", "0")
	res.Header().Set("Transfer-Encoding", "chunked")
	// res.WriteHeader(200)
	fmt.Fprint(res, "hello world")
}

func main() {
	http.HandleFunc("/hello", HandleFunc)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
