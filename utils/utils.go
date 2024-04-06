package utils

import (
	"fmt"
	"net"
)

type TCPServer struct {
	listener net.Listener
	shutdown chan struct{}
	handler  func(net.Conn)
}

func NewTCPServer(address string) (*TCPServer, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on address %s: %w", address, err)
	}

	return &TCPServer{
		listener: listener,
		shutdown: make(chan struct{}),
	}, nil
}

func (s *TCPServer) Start() {
	// TODO select on shutdown channel
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("unable to accept connection: ", err)
			continue
		}
		go s.handler(conn)
	}
}

func (s *TCPServer) Stop() {
	// TODO send msg to shutdown channel
}

const LISTENADDRESS string = "0.0.0.0:4242"
