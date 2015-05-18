package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// http.Error(w, "something failed", http.StatusInternalServerError)
		log.Println("handler")
	}

	req, err := http.NewRequest("GET", "http://godoc.golangtc.com/", nil)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler(w, req)

	fmt.Printf("%d - %s", w.Code, w.Body.String())
}
