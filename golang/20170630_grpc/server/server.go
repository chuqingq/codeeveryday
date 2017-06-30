package main

// server.go

import (
	"log"
	"net"

	pb "../protocol"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type Message struct {
	content    string
	appId      int32
	appName    string
	expireTime int64
	sendTime   int64
	requestId  int32
	callback   string
	deviceId   string
}
type server struct{}
type MessageList []*Message

var devs map[string]MessageList

func (s *server) Push(ctx context.Context, in *pb.PushRequest) (*pb.PushReply, error) {
	//	log.Printf("receive data is %+v\n", in)
	//	now := time.Now().Unix() * 1000 // ms
	//	for _, dev := range in.DeviceIds {
	//		msg := &Message{
	//			content:    in.Message,
	//			appId:      in.AppId,
	//			appName:    in.AppName,
	//			expireTime: in.Delay + in.Expire + now,
	//			sendTime:   in.Delay + now,
	//			requestId:  in.RequestId,
	//			deviceId:   dev,
	//			callback:   "TODO",
	//		}
	//		addMsg(msg)
	//	}
	return &pb.PushReply{Result: 123}, nil
}

func addMsg(msg *Message) {
	_, ok := devs[msg.deviceId]
	if !ok {
		devs[msg.deviceId] = make([]*Message, 0)
	}
	devs[msg.deviceId] = append(devs[msg.deviceId], msg)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	devs = make(map[string]MessageList, 1024)
	s := grpc.NewServer()
	pb.RegisterLogicServer(s, &server{})
	s.Serve(lis)
}
