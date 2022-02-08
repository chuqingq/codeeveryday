package mdns_rpc_sample

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"

	"github.com/hashicorp/mdns"
)

type RPC struct {
	servicesMap sync.Map
}

func NewRPC() *RPC {
	return &RPC{}
}

func (r *RPC) Call(service, method string, args interface{}, reply interface{}) error {
	c, err := r.GetServiceClient(service)
	if err != nil {
		return err
	}
	log.Printf("enter rpc.Call")
	return c.Call(method, args, reply)
}

// TODO func RPCGo(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call {}

type serviceInfo struct {
	ip   string
	port int
	conn *rpc.Client
	lock sync.Mutex // TODO
}

func (r *RPC) GetServiceClient(service string) (*rpc.Client, error) {
	s, ok := r.servicesMap.Load(service)
	if ok {
		return s.(*serviceInfo).conn, nil
	}
	// 如果不存在
	// Make a channel for results and start listening
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	defer close(entriesCh)

	// Start the lookup
	mdns.Lookup(service, entriesCh)
	for entry := range entriesCh {
		fmt.Printf("Got new entry: %v\n", entry)
		// TODO 维护全局services
		addr := fmt.Sprintf("%s:%v", entry.AddrV4, entry.Port)
		log.Printf("addr: %v", addr)
		c, err := rpc.DialHTTP("tcp", addr)
		if err != nil {
			return nil, err
		}
		log.Printf("rpc.Dial success")
		r.servicesMap.Store(service, &serviceInfo{
			ip:   entry.AddrV4.String(),
			port: entry.Port,
			conn: c,
		})
		return c, nil
	}
	return nil, errors.New("service not found")
}

var serversMap sync.Map

// TODO 需要server main中调用Register方法注册到http
func (r *RPC) RegisterService(serviceName string, rcvr interface{}) error {
	port := 8081
	// rpc
	rpc.Register(rcvr)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if e != nil {
		log.Fatalf("listen error: %v", e)
	}
	go http.Serve(l, nil) // TODO

	host, _ := os.Hostname()
	info := []string{"My awesome service"}
	service, err := mdns.NewMDNSService(host, serviceName, "", "", port, nil, info)
	if err != nil {
		log.Printf("NewMDNSService error: %v", err)
		return err
	}
	// Create the mDNS server, defer shutdown
	server, err := mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		log.Printf("NewServer error: %v", err)
		return err
	}
	serversMap.Store(serviceName, server)
	return nil
}

func (r *RPC) UnRegisterService(service string) {
	s, ok := r.servicesMap.Load(service)
	if ok {
		s.(*mdns.Server).Shutdown()
	}
}

func (r *RPC) Close() error {
	// TODO
	return nil
}
