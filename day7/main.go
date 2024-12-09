package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
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

func isValid(target int, ns []int, isPart2 bool) bool {
	if len(ns) == 1 {
		return ns[0] == target
	}
	sum := ns[0] + ns[1]
	if isValid(target, append([]int{sum}, ns[2:]...), isPart2) {
		return true
	}
	prod := ns[0] * ns[1]
	if isValid(target, append([]int{prod}, ns[2:]...), isPart2) {
		return true
	}
	concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", ns[0], ns[1]))
	if isPart2 && isValid(target, append([]int{concat}, ns[2:]...), isPart2) {
		return true
	}
	return false
}

func main() {
	debug = flag.Bool("debug", false, "Debug")
	flag.Parse()

	sum1 := 0
	sum2 := 0
	input := strings.Split(string(input), "\n")
	for _, line := range input {
		parts := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parts[0])
		nsStr := strings.Fields(parts[1])
		ns := make([]int, len(nsStr))
		for i, s := range nsStr {
			ns[i], _ = strconv.Atoi(s)
		}

		if isValid(target, ns, false) {
			sum1 += target
		}
		if isValid(target, ns, true) {
			sum2 += target
		}
	}

	fmt.Println("Part 1: ", sum1)

	fmt.Println("Part 2: ", sum2)
}
