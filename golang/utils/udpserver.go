package utils

import (
	"log"
	"net"
)

type UDPServer struct {
	handler  func(net.Conn)
	shutdown chan struct{}
	addr     *net.UDPAddr
}

func NewUDPServer(address string, handler func(net.Conn)) (*UDPServer, error) {
	addr, err := net.ResolveUDPAddr("udp", LISTENADDRESS)
	if err != nil {
		log.Fatal(err)
	}
	return &UDPServer{
		shutdown: make(chan struct{}),
		handler:  handler,
		addr:     addr,
	}, nil
}

func (s *UDPServer) Start() {
	conn, err := net.ListenUDP("udp", s.addr)
	if err != nil {
		log.Fatal(err)
	}
	s.handler(conn)
}

func (s *UDPServer) Stop() {
	close(s.shutdown)
}
