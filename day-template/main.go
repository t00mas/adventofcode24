package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

var debug *bool

func p(a ...interface{}) {
	if *debug {
		fmt.Println(a...)
	}
}

//go:embed input.txt
var input []byte

func main() {
	debug = flag.Bool("debug", false, "Debug")
	flag.Parse()

	input := strings.Split(string(input), "\n")

	fmt.Println("Part 1: ")

	fmt.Println("Part 2: ")
}
