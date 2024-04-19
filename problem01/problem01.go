package problem01

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"net"
	"os"
	"os/signal"
	"protohackers/utils"
	"syscall"
)

type Request struct {
	Method *string  `json:"method"`
	Number *float64 `json:"number"`
}

type Response struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}

func Run() {
	server, err := utils.NewTCPServer(utils.LISTENADDRESS, handleConnection01)
	if err != nil {
		fmt.Println("error starting server: ", err)
		return
	}
	server.Start()
	// Wait for a SIGINT or SIGTERM signal to gracefully shut down the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("Shutting down server...")
	server.Stop()
	fmt.Println("Server stopped.")
}

func handleConnection01(conn net.Conn) {
	defer conn.Close()
	buf := bufio.NewReader(conn)
	for {
		bytes, err := buf.ReadBytes('\n')
		if err != nil {
			fmt.Println(fmt.Errorf("could not read data: %w", err))
			break
		}
		response := verifyRequest(bytes)
		resp, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Errore sucando json: ", err)
			continue
		}
		resp = append(resp, '\n')
		conn.Write(resp)
		fmt.Println("Sending:", string(resp))
	}
}

func verifyRequest(data []byte) Response {
	var req Request
	fmt.Println("Received:", string(data))
	err := json.Unmarshal(data, &req)
	if err != nil || req.Method == nil || req.Number == nil || *req.Method != "isPrime" {
		return Response{"invalid", false}
	}
	return Response{"isPrime", isPrime(*req.Number)}
}

func isPrime(n float64) bool {
	intn := int(n)
	if intn <= 1 {
		return false
	} else if intn == 2 {
		return true
	} else if intn%2 == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if intn%i == 0 {
			return false
		}
	}
	return true
}

/*
To keep costs down, a hot new government department is contracting out its mission-critical primality testing to the lowest bidder. (That's you).

Officials have devised a JSON-based request-response protocol. Each request is a single line containing a JSON object, terminated by a newline character ('\n', or ASCII 10). Each request begets a response, which is also a single line containing a JSON object, terminated by a newline character.

After connecting, a client may send multiple requests in a single session. Each request should be handled in order.

A conforming request object has the required field method, which must always contain the string "isPrime", and the required field number, which must contain a number. Any JSON number is a valid number, including floating-point values.

Example request:

{"method":"isPrime","number":123}
A request is malformed if it is not a well-formed JSON object, if any required field is missing, if the method name is not "isPrime", or if the number value is not a number.

Extraneous fields are to be ignored.

A conforming response object has the required field method, which must always contain the string "isPrime", and the required field prime, which must contain a boolean value: true if the number in the request was prime, false if it was not.

Example response:

{"method":"isPrime","prime":false}
A response is malformed if it is not a well-formed JSON object, if any required field is missing, if the method name is not "isPrime", or if the prime value is not a boolean.

A response object is considered incorrect if it is well-formed but has an incorrect prime value. Note that non-integers can not be prime.

Accept TCP connections.

Whenever you receive a conforming request, send back a correct response, and wait for another request.

Whenever you receive a malformed request, send back a single malformed response, and disconnect the client.

Make sure you can handle at least 5 simultaneous clients.
*/
