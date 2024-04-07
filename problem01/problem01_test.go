package problem01

import (
	"net"
	"protohackers/utils"
	"testing"
)

func TestServer(t *testing.T) {
	// Start the server
	server, err := utils.NewTCPServer(utils.LISTENADDRESS, handleConnection01)
	if err != nil {
		t.Fatal(err)
	}
	server.Start()
	defer server.Stop()
	conn, err := net.Dial("tcp", utils.LISTENADDRESS)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	request := `{"method":"isPrime","number":123}`
	expected := `{"method":"isPrime","prime":false}`
	actual := make([]byte, 1024)
	if _, err := conn.Write([]byte(request)); err != nil {
		t.Fatal(err)
	}
	if _, err := conn.Read(actual); err != nil {
		t.Fatal(err)
	}
	if string(actual) != expected {
		t.Errorf("expected %q, but got %q", expected, actual)
	}
}
