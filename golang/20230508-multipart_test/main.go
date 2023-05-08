package main

import (
	"bytes"
	"io"
	"log"
	"net/textproto"
	"regexp"
)

type MultipartReader struct {
	reader    io.Reader
	boundary  []byte
	buffer    []byte
	bufferlen int
}

func NewMultipartReader(reader io.Reader, boundary []byte) *MultipartReader {
	return &MultipartReader{
		reader:   reader,
		boundary: boundary,
		buffer:   make([]byte, 0, 1024*1024*16),
	}
}

func (mr *MultipartReader) NextPart() (*Part, error) {
	for {
		// 先从缓冲中尝试获取part
		p := mr.tryGetPart()
		if p != nil {
			return p, nil
		}
		// 如果失败，则继续read
		n, err := mr.reader.Read(mr.buffer[mr.bufferlen:cap(mr.buffer)])
		if err != nil {
			return nil, err
		}
		mr.bufferlen += n
	}
}

var reContentLength = regexp.MustCompile(`Content-Length:.*\r\n`)

// 尝试获取一个part，如果失败，返回nil
func (mr *MultipartReader) tryGetPart() *Part {
	// log.Printf("tryGetPart: bufferlen: %v", mr.bufferlen)
	for mr.bufferlen > 0 {
		ind := bytes.Index(mr.buffer[:mr.bufferlen], mr.boundary)
		// log.Printf("ind: %v", ind)
		if ind < 0 {
			return nil
		}
		body := make([]byte, ind)
		copy(body[:], mr.buffer[:ind])
		body = bytes.Trim(body, "\r\n-")
		// Content-Length
		contentlength := reContentLength.Find(body)
		if len(contentlength) > 0 {
			log.Printf("%s", contentlength)
		}
		// headers
		headerbody := bytes.SplitN(body, []byte("\r\n\r\n"), 2)
		if len(headerbody) == 2 {
			log.Printf("headers: %v", string(headerbody[0]))
			log.Printf("body: %v", string(headerbody[1]))
		}

		mr.bufferlen = copy(mr.buffer[:cap(mr.buffer)], mr.buffer[ind+len(mr.boundary):mr.bufferlen])
		// log.Printf("bufferlen: %v", mr.bufferlen)
		// if true /*mr.bufferlen > 10*/ {
		// 	log.Printf("buffer: %v...", string(mr.buffer[:10]))
		// }

		if len(body) == 0 {
			continue
		}

		return &Part{
			// TODO Header
			Body: body,
		}
	}
	return nil
}

type Part struct {
	// TODO
	Header textproto.MIMEHeader
	// textproto.NewReader().ReadMIMEHeader()
	Body []byte
}

func main() {
	var boundary = []byte("-----------------------7e13971310878")
	content := MockFromFile()
	// ioutil.WriteFile("content.txt", content, 0666)
	mr := NewMultipartReader(bytes.NewReader(content), boundary)
	count := 0
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			log.Printf("EOF")
			return
		}
		if err != nil {
			log.Fatalf("err: %v", err)
		}
		slurp := p.Body
		count += 1
		log.Printf("Part[%v]: Foo:%q, bodysize: %v\n", count, p.Header.Get("Foo"), len(slurp))
	}
}
