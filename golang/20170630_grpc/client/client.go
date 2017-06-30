package main

//client.go

import (
	"log"
	"net/http"
	"strconv"

	pb "../protocol"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

//var conns [100]*grpc.ClientConn
var conn *grpc.ClientConn

func main() {
	//	for i := 0; i < 100; i++ {
	//		conn, err := grpc.Dial(address, grpc.WithInsecure())
	//		if err != nil {
	//			log.Fatalf("did not connect: %v", err)
	//		}
	//		conns[i] = conn
	//	}
	var err error
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	http.HandleFunc("/", testHandler)
	http.ListenAndServe(":55555", nil)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("123"))
	return
	//	conn := conns[rand.Intn(100)]

	//	defer conn.Close()
	c := pb.NewLogicClient(conn)
	res, err := c.Push(context.Background(), &pb.PushRequest{AppId: 1, AppName: "xxx", Expire: 123, Delay: 234, RequestId: 12, Message: "test", DeviceIds: []string{"12334", "234235"}})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//	log.Printf("result: %d\n", r.Result)
	w.Write([]byte(strconv.Itoa(int(res.Result))))
}
