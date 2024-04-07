package utils

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type TCPServer struct {
	wg         sync.WaitGroup
	listener   net.Listener
	shutdown   chan struct{}
	handler    func(net.Conn)
	connection chan net.Conn
}

func NewTCPServer(address string, handler func(net.Conn)) (*TCPServer, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on address %s: %w", address, err)
	}

	return &TCPServer{
		listener:   listener,
		shutdown:   make(chan struct{}),
		connection: make(chan net.Conn),
		handler:    handler,
	}, nil
}

func (s *TCPServer) acceptConnections() {
	defer s.wg.Done()
	for {
		select {
		case <-s.shutdown:
			return
		default:
			conn, err := s.listener.Accept()
			if err != nil {
				continue
			}
			s.connection <- conn
		}
	}
}

func (s *TCPServer) Start() {
	s.wg.Add(2)
	go s.acceptConnections()
	go s.handleConnections()
}

func (s *TCPServer) handleConnections() {
	defer s.wg.Done()
	// TODO select on shutdown channel
	for {
		select {
		case <-s.shutdown:
			return
		case conn := <-s.connection:
			go s.handler(conn)
		}
	}
}

func (s *TCPServer) Stop() {
	close(s.shutdown)
	s.listener.Close()
	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()
	select {
	case <-done:
		return
	case <-time.After(time.Second):
		fmt.Println("Timed out waiting for connections to finish.")
		return
	}
}

const LISTENADDRESS string = "0.0.0.0:4242"
