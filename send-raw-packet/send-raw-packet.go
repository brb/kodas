// send raw (IPV4) packet
package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"syscall"
)

func exitOnErr(msg string, err error) {
	if err != nil {
		exit(msg)
	}
}

func exit(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(-1)
}

func main() {
	if len(os.Args) != 3 {
		exit(fmt.Sprintf("usage: %s <iface> <packet-in-hex-str>", os.Args[0]))
	}

	packet, err := hex.DecodeString(os.Args[2])
	exitOnErr(fmt.Sprintf("cannot decode hex str: %s", err), err)

	proto := 0x0008 // htons(ETH_P_IPV4)
	sock, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, proto)
	exitOnErr(fmt.Sprintf("cannot create socket: %s", err), err)
	defer func() { syscall.Close(sock) }()

	ifaceName := os.Args[1]
	iface, err := net.InterfaceByName(ifaceName)
	exitOnErr(fmt.Sprintf("cannot find %q interface: %s", ifaceName, err), err)

	dst := syscall.SockaddrLinklayer{Ifindex: iface.Index}
	err = syscall.Sendto(sock, packet, 0, &dst)

	exitOnErr(fmt.Sprintf("cannot send: %s", err), err)
}
