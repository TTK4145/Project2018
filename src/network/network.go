package network

import (
	"fmt"
	"log"
	"net"
)

const max_udp_datagram_size = 508

func SendBytes(data []byte, destination string) error {
	if len(data) >= max_udp_datagram_size {
		return fmt.Errorf("network: too many bytes of data (%i)", len(data))
	}

	addr, err := net.ResolveUDPAddr("udp", destination)
	if err != nil {
		return fmt.Errorf("network.SendBytes: %s", err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return fmt.Errorf("network.SendBytes: %s", err)
	}
	defer conn.Close()

	n, err := conn.Write(data)
	if err != nil {
		return fmt.Errorf("network.SendBytes: %s", err)
	}

	log.Printf("%d bytes of data sent too %s\n", n, destination)


	return nil
}
