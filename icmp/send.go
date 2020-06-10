package icmp

import (
	"encoding/binary"
	"fmt"
	"golang.org/x/net/ipv4"
	"log"
	"net"
	"syscall"
	"time"
)

func Send(destinationIp string, timeChan chan time.Time) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		log.Println(err)
	}
	addr := syscall.SockaddrInet4{
		Port: 0,
		Addr: [4]byte{127, 0, 0, 1},
	}

	err = syscall.SetsockoptInt(fd, syscall.IPPROTO_IP, syscall.IP_HDRINCL, 1)
	if err != nil {
		panic(err)
	}

	destIp := net.ParseIP(destinationIp)
	data := makepacket(destIp)

	//for _, v := range data {
	//	if v == 0 {
	//		fmt.Printf("00 ")
	//		continue
	//	} else if v < 0xf {
	//		fmt.Printf("0%x ", v)
	//		continue
	//	}
	//	fmt.Printf("%x ", v)
	//}
	//fmt.Printf("\n")


	fmt.Printf("[pingo]: (send) ==> %v\n", destIp)
	go func() {
		timeChan <- time.Now()
	}()
	err = syscall.Sendto(fd, data, 0, &addr)

	if err != nil {
		panic(err)
	}
}

func makepacket(destinationIP net.IP) []byte {
	icmp := []byte{
		8, // type: echo request
		0, // code: not used by echo request
		0, // checksum (16 bit), we fill in below
		0,
		0, // identifier (16 bit). zero allowed.
		0,
		0, // sequence number (16 bit). zero allowed.
		0,
	}
	//icmp = append(icmp, []byte(s)...)

	cs := csum(icmp)
	icmp[2] = byte(cs)
	icmp[3] = byte(cs >> 8)

	h := &ipv4.Header{
		Version:  ipv4.Version,
		TOS: 0,
		Len:      ipv4.HeaderLen,
		TotalLen: ipv4.HeaderLen + len(icmp), // 20 bytes for IP, 10 for ICMP
		TTL:      64,
		Protocol: 1, // ICMP
		Dst:      destinationIP,
		// ID, Src and Checksum will be set for us by the kernel
	}

	buf, err := h.Marshal()
	if err != nil {
		log.Println(err)
	}

	binary.LittleEndian.PutUint16(buf[2:4], uint16(len(icmp) + len(buf)))
	return append(buf, icmp...)
}

func csum(b []byte) uint16 {
	var s uint32
	for i := 0; i < len(b); i += 2 {
		s += uint32(b[i+1])<<8 | uint32(b[i])
	}
	// add back the carry
	s = s>>16 + s&0xffff
	s = s + s>>16
	return uint16(^s)
}
