package main

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"strings"
)

var str = "---------------------------7e13971310878\r\nFoo: one\r\n\r\nA section1123123123123123123\r\n" +
	"---------------------------7e13971310878\r\nFoo: two\r\n\r\nAnd another\r\n" +
	"---------------------------7e13971310878\r\n"

func main() {
	mr := multipart.NewReader(strings.NewReader(str), "-------------------------7e13971310878")
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
		fmt.Printf("Part %q: %q\n", p.Header.Get("Foo"), slurp)
	}
}

