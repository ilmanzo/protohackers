package problem02

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
)

type PriceItem struct {
	timestamp int32
	price     int32
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
	var pricehistory []PriceItem
	in_msgbuf := make([]byte, 9)
	out_msgbuf := make([]byte, 4)
	for {
		bytes, err := io.ReadFull(conn, in_msgbuf)
		if bytes < 9 || err != nil {
			fmt.Println(fmt.Errorf("could not read data: %w", err))
			break
		}
		a := int32(binary.BigEndian.Uint32(in_msgbuf[1:5]))
		b := int32(binary.BigEndian.Uint32(in_msgbuf[5:]))
		switch in_msgbuf[0] {
		case 'I':
			pricehistory = append(pricehistory, PriceItem{a, b})
			//fmt.Printf("history: %v\n", pricehistory)
		case 'Q':
			mean := calc_mean(pricehistory, a, b)
			//fmt.Printf("query: %v - %v -> mean %v \n", a, b, mean)
			binary.BigEndian.PutUint32(out_msgbuf, uint32(mean))
			conn.Write(out_msgbuf)
		}
	}
}

func calc_mean(pricehistory []PriceItem, time_start int32, time_end int32) int32 {
	var total int64
	var n int64
	for _, item := range pricehistory {
		//fmt.Printf("item timestamp: %v\n", item.timestamp)
		if item.timestamp >= time_start && item.timestamp <= time_end {
			total += int64(item.price)
			n += 1
		}
	}
	if n == 0 {
		return 0
	}
	return int32(total / n)
}
