package main

import (
	"flag"
	"fmt"
	"os"
	"protohackers/problem00"
	"protohackers/problem01"
	"protohackers/problem02"
	"protohackers/problem03"
	"protohackers/problem04"
	"protohackers/utils"
)

func main() {
	problems := []func(){
		problem00.Run,
		problem01.Run,
		problem02.Run,
		problem03.Run,
		problem04.Run}
	problem := flag.Int("problem", -1, "the problem to run")
	flag.Parse()
	if *problem < 0 || *problem > (len(problems)-1) {
		fmt.Println("You want problem = ", *problem)
		fmt.Println("Please specify a problem between 0 and", len(problems)-1)
		os.Exit(1)
	}
	fmt.Printf("Running problem %v, listening on %v\n", *problem, utils.TCP_LISTENADDRESS)
	problems[*problem]()
}
