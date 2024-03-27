package main

import (
	"flag"
	"fmt"
	"protohackers/problem00"
	"protohackers/problem01"
)

func main() {
	problem := flag.Int("problem", 0, "the problem to run")
	listenaddress := "0.0.0.0:4242"
	flag.Parse()
	fmt.Printf("Running problem %v, listening on %v\n", *problem, listenaddress)
	switch *problem {
	case 0:
		problem00.Run(listenaddress)
	case 1:
		problem01.Run(listenaddress)
	default:
		fmt.Println("Not yet implemented")
	}
}
