/* Prints CIDRs which cover the given [from; to] IP range */
package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	fullMask         = ^uint32(0)
	cidrMaxPrefixLen = 32
)

type CIDR struct {
	ip       net.IP
	maskSize int
}

type CIDRs []CIDR

func (cidr CIDR) String() string {
	return fmt.Sprintf("%s/%d", cidr.ip, cidr.maskSize)
}

func (cidrs CIDRs) String() string {
	l := make([]string, len(cidrs))

	for i, c := range cidrs {
		l[i] = fmt.Sprintf("%s", c)
	}

	return strings.Join(l, "\n")
}

func parseCIDRs(from, to net.IP) (cidrs CIDRs) {
	var (
		start = ipToUint32(from)
		end   = ipToUint32(to)
	)

	for end >= start {
		mask, maskSize := fullMask, cidrMaxPrefixLen
		for mask > 0 {
			tmpMask := mask << 1
			if (start&tmpMask) != start || (start|^tmpMask) > end {
				break
			}
			mask = tmpMask
			maskSize--
		}
		cidrs = append(cidrs, CIDR{uint32ToIP(start), maskSize})
		start |= ^mask
		if start+1 < start {
			break
		}
		start++
	}

	return cidrs
}

func parseIPOrFail(ip string) net.IP {
	pip := net.ParseIP(ip)
	if pip == nil {
		fmt.Fprintf(os.Stderr, "invalid ip: %s", ip)
		os.Exit(-2)
	}
	return pip
}

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
	fmt.Printf("usage: %s <from-ip> <to-ip>\n", prog)
}

func main() {
	if len(os.Args) != 3 {
		usage(os.Args[0])
		os.Exit(-1)
	}
	from := parseIPOrFail(os.Args[1])
	to := parseIPOrFail(os.Args[2])

	fmt.Println(parseCIDRs(from, to))
}
