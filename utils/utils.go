package utils

import (
	"fmt"
	"net"
	"os"
)

const LISTENADDRESS string = "0.0.0.0:4242"

func NewTCPListener(listenaddress string) net.Listener {
	listener, err := net.Listen("tcp", listenaddress)
	if err != nil {
		fmt.Println("unable to create listener: ", err)
		os.Exit(1)
	}
	return listener
}

func NewUDPListener(listenaddress string) net.Listener {
	addr := &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 4242,
		Zone: "",
	}
	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("unable to create listener: ", err)
		os.Exit(1)
	}
	return listener
}
