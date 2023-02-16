package udp_utils

import (
	"fmt"
	"net"
)

func SendPacket(remoteAddr string, portUDP int, data []byte) {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP(remoteAddr),
		Port: portUDP,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
