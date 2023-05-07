package main

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com/")
	// resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
	if err != nil {
		log.Fatalf("http req error: %v", err)
	}
	defer resp.Body.Close()

	mr := multipart.NewReader(resp.Body, "-------------------------7e13971310878")
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("err: %v", err)
		}
		slurp, err := io.ReadAll(p)
		if err != nil {
			log.Fatalf("read: %v", err)
		}
		log.Printf("Part %q: %q\n", p.Header.Get("Foo"), slurp)
	}
}

