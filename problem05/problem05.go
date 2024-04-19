package problem05

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"protohackers/utils"
	"regexp"
	"strings"
	"syscall"
)

func Run() {
	server, err := utils.NewTCPServer(utils.LISTENADDRESS, handleConnection)
	if err != nil {
		fmt.Println("error starting server: ", err)
		return
	}
	server.Start()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	fmt.Println("Shutting down server...")
	server.Stop()
}

var boguscoin = regexp.MustCompile(`^7[a-zA-Z0-9]{25,34}$`)

func relay(dst io.WriteCloser, src io.ReadCloser) {
	defer func() { src.Close(); dst.Close() }()

	for r := bufio.NewReader(src); ; {
		msg, err := r.ReadString('\n')
		if err != nil {
			return
		}

		tokens := make([]string, 0, 8)
		for _, raw := range strings.Split(msg[:len(msg)-1], " ") {
			t := boguscoin.ReplaceAllString(raw, "7YWHMfk9JZe0LM0g1ZauHuiSxhI")
			tokens = append(tokens, t)
		}

		out := strings.Join(tokens, " ") + "\n"
		if _, err = dst.Write([]byte(out)); err != nil {
			log.Printf("error in writing: %s", err)
		}
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	upstream, err := net.Dial("tcp", "chat.protohackers.com:16963")
	if err != nil {
		log.Printf("cannot connect upstream: %s", err)
		conn.Close()
		return
	}
	go relay(conn, upstream)
	relay(upstream, conn)
}
