package main

import (
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/ghettovoice/gosip/log"
	"github.com/ghettovoice/gosip/sip"
	"github.com/ghettovoice/gosip/sip/parser"
)

var logger *log.LogrusLogger

func main() {
	logger = log.NewDefaultLogrusLogger()
	logger.SetLevel(log.DebugLevel)

	// tcp server listen at :9000
	startTCPServer()

	// sip server
	runSipServer()
}

func runSipServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:5060")
	if err != nil {
		logger.Fatalf("sip server listen error: %v", err)
	}
	logger.Infof("sip server listen success")

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Fatalf("sip server accept error: %v", err)
		}
		logger.Infof("sip server accept: %v", conn.RemoteAddr())

		go handleSipConn(conn)
	}
}

func handleSipConn(conn net.Conn) {
	defer conn.Close()

	// parser
	msgChan := make(chan sip.Message, 8)
	errChan := make(chan error, 8)
	parser := parser.NewParser(msgChan, errChan, true, logger)

	go io.Copy(parser, conn)

	// handle sip
	for {
		select {
		case msg := <-msgChan:
			logger.Debugf("recv msg: %v", msg.String())
			handleSip(msg, conn)
		case err := <-errChan:
			logger.Errorf("recv err: %v", err)
			// return
		}
	}
}

var inviteReq sip.Request

func handleSip(msg sip.Message, conn net.Conn) {
	switch msg := msg.(type) {
	case sip.Request:
		// 请求
		switch msg.Method() {
		case sip.REGISTER:
			// 回复200 OK
			resp := sip.NewResponseFromRequest(msg.MessageID(), msg, 200, "OK", "")
			conn.Write([]byte(resp.String()))
			logger.Debugf("send register resp: %v", resp.String())
			// 发送invite
			inviteReq = invite()
			conn.Write([]byte(inviteReq.String()))
			logger.Debugf("send invite req: %v", inviteReq.String())
		case sip.MESSAGE:
			// 回复 200 OK
			resp := sip.NewResponseFromRequest(msg.MessageID(), msg, 200, "OK", "")
			conn.Write([]byte(resp.String()))
			logger.Debugf("send message resp: %v", resp.String())

		}
	case sip.Response:
		// 一般响应无需理会
		cseq, ok := msg.CSeq()
		if ok && strings.Contains(cseq.String(), string(sip.INVITE)) && msg.StatusCode() == 200 {
			// 如果是INVITE的200 OK，则需要回复ACK
			ack := sip.NewAckRequest(msg.MessageID(), inviteReq, msg, "", nil)
			conn.Write([]byte(ack.String()))
			logger.Debugf("send invite ack: %v", ack.String())
		}
	}
}

func invite() sip.Request {
	// 发送invite
	rawMsg := []string{
		"INVITE sip:34020000001320000001@3402000000 SIP/2.0",
		"Via: SIP/2.0/TCP 192.168.0.140:5060;rport;branch=z9hG4bK36029g",
		"From: <sip:34020000002000000001@3402000000>;tag=SRS3k96w5t6",
		"To: <sip:34020000001320000001@3402000000>",
		"CSeq: 706 INVITE",
		"Call-ID: 092q635330745o4f",
		"User-Agent: SRS/6.0.48(Bee)",
		"Contact: <sip:34020000002000000001@3402000000>",
		"Subject: 34020000001320000001:0200004778,34020000002000000001:0",
		"Max-Forwards: 70",
		"Content-Type: Application/SDP",
		"Content-Length: 170",
		"Connection: Keep-Alive",
		"",
		"v=0",
		"o=34020000001320000001 0 0 IN IP4 192.168.0.140",
		"s=Play",
		"c=IN IP4 192.168.0.140",
		"t=0 0",
		"m=video 9000 TCP/RTP/AVP 96",
		"a=recvonly",
		"a=rtpmap:96 PS/90000",
		"y=0200004778",
		"",
	}
	req := Message(rawMsg, logger)
	return req
}

func Message(rawMsg []string, logger *log.LogrusLogger) sip.Request {
	msg, err := parser.ParseMessage([]byte(strings.Join(rawMsg, "\r\n")), logger)
	if err != nil {
		logger.Errorf("parse err: %v", err)
	}
	// Expect(err).ToNot(HaveOccurred())
	switch msg := msg.(type) {
	case sip.Request:
		return msg
	case sip.Response:
		panic(fmt.Sprintf("%s is not a request", msg.Short()))
	default:
		panic(fmt.Sprintf("%s is not a request", msg.Short()))
	}
}

func startTCPServer() {
	// 监听tcp的9000端口，并接收消息
	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		panic("error listening:" + err.Error())
	}
	logger.Infof("Starting the server")

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				panic("Error accept:" + err.Error())
			}
			logger.Infof("Accepted the Connection: %v", conn.RemoteAddr())
			go HandleTCPConn(conn)
		}
	}()
}

func HandleTCPConn(conn net.Conn) {
	buf := make([]byte, 10240)
	defer conn.Close()

	for {
		n, err := conn.Read(buf)
		switch err {
		case nil:
			logger.Infof("tcpserver read %v bytes\n", n)
		case io.EOF: // notice
			logger.Errorf("Warning: End of data: %s \n", err)
			return
		default:
			logger.Errorf("Error: Reading data : %s \n", err)
			return
		}
	}
}
