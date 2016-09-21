// Dump ethernet frame
package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <packet-in-hexstr>\n", os.Args[0])
		os.Exit(-1)
	}

	packet, err := hex.DecodeString(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot decode packet in hex: %s\n", err)
		os.Exit(-1)
	}

	fmt.Println(gopacket.NewPacket(packet, layers.LayerTypeEthernet, gopacket.Default).Dump())
}
