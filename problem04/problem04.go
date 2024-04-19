package problem04

import (
	"fmt"
	"log"
	"net"
	"protohackers/utils"
	"strings"
)

func Run() {
	server, err := utils.NewUDPServer(utils.LISTENADDRESS, handleConnection)
	if err != nil {
		fmt.Println("error starting server: ", err)
		return
	}
	server.Start()
	server.Stop()
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	udpConn, ok := conn.(*net.UDPConn)
	if !ok {
		log.Println("bad connection type")
		return
	}

	var store = map[string]string{"version": "Ken's Key-Value Store 1.0.0"}

	buf := make([]byte, 1024)
	for {
		n, addr, err := udpConn.ReadFrom(buf)
		if err != nil {
			log.Print(err)
			continue
		}

		request := string(buf[:n])
		log.Printf("%v %v", addr, request)

		k, v, insert := strings.Cut(request, "=")
		if insert {
			if k == "version" {
				continue
			}
			store[k] = v

		} else {
			response := fmt.Sprintf("%v=%v", k, store[k])

			_, err := udpConn.WriteTo([]byte(response), addr)
			if err != nil {
				log.Print(err)
			}
		}
	}
}
