/* Converts genmask (e.g. 255.255.240.0) to CIDR prefix */
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

func usage(prog string) {
	fmt.Printf("usage: %s <cidr>\n", prog)
}

func main() {
	if len(os.Args) != 2 {
		usage(os.Args[0])
		os.Exit(-1)
	}

	ip := ipToUint32(net.ParseIP(os.Args[1]))
	i := 0
	for ; ip != 0; i++ {
		ip <<= 1
	}

	fmt.Printf("/%d\n", i)
}
