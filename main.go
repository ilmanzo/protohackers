package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"protohackers/problem00"
	"protohackers/problem01"
	"protohackers/problem02"
	"protohackers/problem03"
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

func main() {
	handlers := []func(conn net.Conn){
		problem00.HandleConnection,
		problem01.HandleConnection,
		problem02.HandleConnection,
		problem03.HandleConnection}
	problem := flag.Int("problem", -1, "the problem to run")
	flag.Parse()
	if *problem < 0 || *problem > (len(handlers)-1) {
		fmt.Println("You want problem = ", *problem)
		fmt.Println("Please specify a problem between 0 and", len(handlers)-1)
		os.Exit(1)
	}
	fmt.Printf("Running problem %v, listening on %v\n", *problem, LISTENADDRESS)
	listener := NewTCPListener(LISTENADDRESS)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("unable to accept connection: ", err)
			continue
		}
		go handlers[*problem](conn)
	}

}
