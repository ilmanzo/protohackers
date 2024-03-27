package problem01

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Request struct {
	Method string `json:"method"`
	Number int    `json:"number"`
}

type Response struct {
	Method string `json:"method"`
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
	var req Request
	json.Unmarshal(buffer[:n], &req)
	fmt.Println("Received: ", string(buffer[:n]))
	resp := Response{Method: "isPrime", Prime: isPrime(req.Number)}
	r, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error serializing json: ", err)
		return
	}
	fmt.Println("Sending: ", string(r))
	conn.Write(r)
	conn.Write([]byte{10})
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
