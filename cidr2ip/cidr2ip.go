/* Expands CIDR */
package main

import (
	"fmt"
	"net"
	"os"
)

func ipToUint32(ip net.IP) (r uint32) {
	for _, b := range ip.To4() {
		r <<= 8
		r |= uint32(b)
	}

	return
}

func uint32ToIP(ip uint32) net.IP {
	return net.IPv4(
		byte(ip>>24),
		byte((ip>>16)&0xff),
		byte((ip>>8)&0xff),
		byte(ip&0xff),
	)
}

func usage(prog string) {
	fmt.Printf("usage: %s <cidr>\n", prog)
}

func main() {
	if len(os.Args) != 2 {
		usage(os.Args[0])
		os.Exit(-1)
	}
	ip, ipnet, err := net.ParseCIDR(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid cidr %s: %s", os.Args[1], err)
	}

	ones, bits := ipnet.Mask.Size()
	fmt.Println(ip)
	fmt.Println(uint32ToIP(ipToUint32(ip) | (1<<uint32(bits-ones) - 1)))
}
