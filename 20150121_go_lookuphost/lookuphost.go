package main

import (
	"log"
	"net"
)

func main() {
	addrs, err := net.LookupHost("localhost")
	log.Printf("LookupHost: %v, %v\n", addrs, err)

	ips, err := net.LookupIP("localhost")
	log.Printf("LookupIP: %v, %v\n", ips, err)

	iaddrs, err := net.InterfaceAddrs()
	log.Printf("InterfaceAddrs: %v, %v\n", iaddrs, err)
}

// 2015/01/21 09:35:48 LookupHost: [::1 127.0.0.1], <nil>
// 2015/01/21 09:35:48 LookupIP: [::1 127.0.0.1], <nil>
// 2015/01/21 09:35:48 InterfaceAddrs: [10.45.2.218 0.0.0.0 192.168.13.158], <nil>
