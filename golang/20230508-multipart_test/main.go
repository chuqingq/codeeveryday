package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/textproto"
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
		buffer:   make([]byte, 0, 1024*1024*4),
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

// 尝试获取一个part，如果失败，返回nil
func (mr *MultipartReader) tryGetPart() *Part {
	for mr.bufferlen > 0 {
		ind := bytes.Index(mr.buffer[:mr.bufferlen], mr.boundary)
		if ind < 0 {
			return nil
		}
		data := append([]byte{}, mr.buffer[:ind]...)
		data = bytes.Trim(data, "\r\n-")
		// header and body
		headerAndBody := bytes.SplitN(data, []byte("\r\n\r\n"), 2)
		var headers textproto.MIMEHeader
		var body []byte
		var err error
		if len(headerAndBody) == 2 {
			// headers
			headers, err = textproto.NewReader(bufio.NewReader(bytes.NewReader(append(headerAndBody[0], []byte("\r\n\r\n")...)))).ReadMIMEHeader()
			if err != nil {
				log.Printf("headers error: %v", err)
			}
			// body
			body = headerAndBody[1]
		}

		// TODO
		mr.bufferlen = copy(mr.buffer[:cap(mr.buffer)], mr.buffer[ind+len(mr.boundary):mr.bufferlen])

		if len(data) == 0 {
			continue
		}
		return &Part{
			Header: headers,
			Body:   body,
		}
	}
	return nil
}

type Part struct {
	Header textproto.MIMEHeader
	Body   []byte
}

func main() {
	var boundary = []byte("-----------------------7e13971310878")
	content := MockFromFile()
	// ioutil.WriteFile("content.txt", content, 0666)

	mr := NewMultipartReader(bytes.NewReader(content), boundary)
	for i := 1; ; i++ {
		p, err := mr.NextPart()
		if err == io.EOF {
			log.Printf("EOF")
			return
		}
		if err != nil {
			log.Fatalf("err: %v", err)
		}

		body := p.Body
		var bodystr = ""
		if len(body) > 6 {
			bodystr = fmt.Sprintf("%v...%v", string(body[:6]), string(body[len(body)-6:]))
		} else {

		}
		log.Printf("Part[%v]: Content-Length:%q, bodysize: %v, body: %v\n", i, p.Header.Get("Content-Length"), len(body), bodystr)
	}
}

// 2023/05/08 20:39:05 Part[36]: Content-Length:"506", bodysize: 504, body: <Event...Alert>
// 2023/05/08 20:39:05 Part[37]: Content-Length:"4123", bodysize: 4121, body: <?xml ...Alert>
// 2023/05/08 20:39:05 Part[38]: Content-Length:"24708", bodysize: 24708, body: �����...x/���
// 2023/05/08 20:39:05 Part[39]: Content-Length:"417908", bodysize: 417908, body: ����j...�g��
// 2023/05/08 20:39:05 EOF
