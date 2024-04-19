package problem00

import (
	"net"
	"protohackers/utils"
	"testing"
)

func TestServer(t *testing.T) {
	// Start the server
	server, err := utils.NewTCPServer(utils.LISTENADDRESS, handleConnection)
	if err != nil {
		t.Fatal(err)
	}
	server.Start()
	defer server.Stop()
	// Connect to the server and send a message
	conn, err := net.Dial("tcp", utils.LISTENADDRESS)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	msg := "Testing the echo server!\n"
	actual := make([]byte, len(msg))
	if _, err := conn.Write([]byte(msg)); err != nil {
		t.Fatal(err)
	}
	if _, err := conn.Read(actual); err != nil {
		t.Fatal(err)
	}
	if string(actual) != msg {
		t.Errorf("expected %q, but got %q", msg, actual)
	}
}
