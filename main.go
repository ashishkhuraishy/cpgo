package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		log.Fatalln("not enough argumnents")
	}

	fmt.Println(args)
}
