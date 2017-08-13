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
	conn, err := net.ListenPacket("ip4:udp", "127.0.0.1")
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
		if err := gopacket.SerializeLayers(buf, opts, res, gopacket.Payload("response\n")); err != nil {
			log.Fatal(err)
		}


		if _, err := conn.WriteTo(buf.Bytes(), ipaddr); err != nil {
			log.Fatal(err)
		}
	}
}

/*
# drop ICMP(destination unreachable) output
sudo iptables -A OUTPUT  -p icmp --icmp-type 3 -j DROP

$ sudo ./rawsock_udp_send
2017/08/13 15:22:32 recv from ip: 127.0.0.1
2017/08/13 15:22:32 recv udp port: 54535 -> 12345(italk)
2017/08/13 15:22:32 recv from ip: 127.0.0.1
2017/08/13 15:22:32 recv udp port: 12345(italk) -> 54535

$ nc -u 127.0.0.1 12345
123
response

*/
