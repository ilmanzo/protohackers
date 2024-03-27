package problem01

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"net"
	"os"
)

type Request struct {
	Method *string `json:"method"`
	Number *int    `json:"number"`
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
	buf := bufio.NewReader(conn)
	for {
		bytes, err := buf.ReadBytes('\n')

		if err != nil {
			fmt.Println(fmt.Errorf("could not read data: %w", err))
			break
		}
		var req Request
		json.Unmarshal(bytes, &req)
		fmt.Println("Received:", string(bytes))
		var response Response
		if req.Method == nil || req.Number == nil || *req.Method != "isPrime" {
			// generate invalid response
			response = Response{"invalid", false}
		} else {
			response = Response{"isPrime", isPrime(*req.Number)}
		}
		resp, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error serializing json: ", err)
			continue
		}
		resp = append(resp, '\n')
		conn.Write(resp)
		fmt.Println("Sending:", string(resp))
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
