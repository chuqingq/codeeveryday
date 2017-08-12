package main

/*
echo server by udp for rawsocket
*/

import (
	"log"
	"net"
	//"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func main() {
	conn, err := net.ListenPacket("ip4:udp", "0.0.0.0")
	if err != nil {
		log.Fatalf("ListenPacket error: %v\n", err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Printf("ReadFrom error: %v\n", err)
		}

		if addr.String() != "127.0.0.1" {
			continue
		}
		log.Printf("recv from ip: %v", addr.String())

		ipaddr, _ := net.ResolveIPAddr(addr.Network(), addr.String())

		reqPacket := gopacket.NewPacket(buf[:n], layers.LayerTypeUDP, gopacket.Default)
		udpLayer := reqPacket.Layer(layers.LayerTypeUDP)
		if udpLayer == nil {
			log.Printf("udpLayer nil")
			continue
		}
		req := udpLayer.(*layers.UDP)

		log.Printf("recv udp port: %v -> %v", req.SrcPort, req.DstPort)
		if req.DstPort != 12345 {
			continue
		}
		
		// Our IP header... not used, but necessary for TCP checksumming.
		ip := &layers.IPv4{
			SrcIP:    ipaddr.IP, // TODO
			DstIP:    ipaddr.IP,
			Protocol: layers.IPProtocolTCP,
		}

		res := &layers.UDP{
			SrcPort: req.DstPort,
			DstPort: req.SrcPort,
		}
		res.SetNetworkLayerForChecksum(ip)

		// response
		buf := gopacket.NewSerializeBuffer()
		opts := gopacket.SerializeOptions{
			ComputeChecksums: true,
			FixLengths:       true,
		}
		if err := gopacket.SerializeLayers(buf, opts, res, gopacket.Payload("response\r")); err != nil {
			log.Fatal(err)
		}


		if _, err := conn.WriteTo(buf.Bytes(), ipaddr); err != nil {
			log.Fatal(err)
		}
		/*
		// 交换源端和目的端端口
		buf[0], buf[1], buf[2], buf[3] = buf[2], buf[3], buf[0], buf[1]
		conn.WriteTo(buf[:n], addr)
		*/
	}
}

/*
$ sudo ./rawsock_udp 
2017/08/11 16:09:24 ReadFrom 127.0.0.1, 24, 123123123123123

$ nc -u 127.0.0.1 12345
123123123123123

*/
