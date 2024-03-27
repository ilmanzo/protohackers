package problem01

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Message struct {
	Method string `json:"method"`
	Number int    `json:"number"`
	Prime  bool   `json:"prime"`
}

func Run(listenaddress string) {
	listener, err := net.Listen("tcp", listenaddress)
	if err != nil {
		fmt.Println("unable to create listener: ", err)
		os.Exit(1)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("unable to accept connection: ", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading: ", err)
		return
	}
	var m Message
	json.Unmarshal(buffer[:n], &m)
	m.Prime = isPrime(m.Number)
	resp, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error serializing json: ", err)
		return
	}
	conn.Write(resp)
}

func isPrime(n int) bool {
	if n == 1 || n == 2 || n == 3 {
		return true
	}
	return false
}
