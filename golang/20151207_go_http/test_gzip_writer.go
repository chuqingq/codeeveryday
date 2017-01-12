package main

/*
验证gzip.Writer.Write()返回的第一个参数为：输入的字节数

chuqq@chuqq-hp:~/tmp$ go run test_gzip_writer.go
write: 10
result: [31 139 8 0 0 9 110 136 0 255 50 52 50 54 49 53 51 183 176 52 0 4 0 0 255 255 229 174 29 38 10 0 0 0]
*/

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

func main() {
	var result bytes.Buffer
	w := gzip.NewWriter(&result)
	n, _ := w.Write([]byte("1234567890"))
	println("write:", n)
	w.Close()
	fmt.Println("result:", result.Bytes())
}
