package problem04

import (
	"bytes"
	"fmt"
	"net"
	"protohackers/utils"
)

func Run() {
	server, err := utils.NewTCPServer(utils.LISTENADDRESS, handleConnection)
	if err != nil {
		fmt.Println("error starting server: ", err)
		return
	}
	server.Start()
	server.Stop()
}

const bufSize int = 65535

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, bufSize)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			if err != nil {
				fmt.Println(fmt.Errorf("could not read data: %w", err))
				break
			}
		}
		if bytes.Equal(buf, []byte("version")) {
			response := "version=ilmanzo's Key-Value Store 1.0"
			conn.Write([]byte(response))
		}
	}
}
