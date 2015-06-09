package main

import (
	"log"
	"net"
)

type WhiteItem struct {
	IP     string
	IPMask net.IPMask
}

func Mask(ipstr string, whiteList []WhiteItem) bool {
	ip := net.ParseIP(ipstr)
	for _, item := range whiteList {
		if ip.Mask(item.IPMask).String() == item.IP {
			return true
		}
	}

	return false
}

func main() {
	whiteList := make([]WhiteItem, 0)
	whiteList = append(whiteList, WhiteItem{IP: "192.168.254.0", IPMask: net.CIDRMask(24, 32)})

	log.Printf("%v\n", net.ParseIP("255.255.254.0").DefaultMask().String())
	log.Printf("%v\n", net.CIDRMask(23, 32).String())

	log.Printf("%v\n", Mask("192.168.255.1", whiteList)) // true
	log.Printf("%v\n", Mask("192.168.254.1", whiteList)) // false
}
