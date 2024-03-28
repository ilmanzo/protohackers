package main

import (
	"flag"
	"fmt"
	"os"
	"protohackers/problem00"
	"protohackers/problem01"
	"protohackers/problem02"
)

const LISTENADDRESS string = "0.0.0.0:4242"

func main() {
	problems := []func(string){problem00.Run, problem01.Run, problem02.Run}
	problem := flag.Int("problem", -1, "the problem to run")
	flag.Parse()
	if *problem < 0 || *problem > (len(problems)-1) {
		fmt.Println("You want problem = ", *problem)
		fmt.Println("Please specify a problem between 0 and", len(problems)-1)
		os.Exit(1)
	}
	fmt.Printf("Running problem %v, listening on %v\n", *problem, LISTENADDRESS)
	problems[*problem](LISTENADDRESS)
}
