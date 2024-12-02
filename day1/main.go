package main

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input_txt []byte

func main() {
	input := string(input_txt)
	n1s := make([]int, 0)
	n2s := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		ns := strings.Split(line, "   ")
		n1, e := strconv.Atoi(ns[0])
		if e != nil {
			panic(e)
		}
		n1s = append(n1s, n1)
		n2, e := strconv.Atoi(ns[1])
		if e != nil {
			panic(e)
		}
		n2s = append(n2s, n2)
	}
	slices.Sort(n1s)
	slices.Sort(n2s)
	sum1 := 0
	for i := 0; i < len(n1s); i++ {
		sum1 += max(n1s[i]-n2s[i], n2s[i]-n1s[i])
	}
	println("Part 1: ", sum1)

	times := make(map[int]int)
	for _, n := range n1s {
		for _, m := range n2s {
			if m < n {
				continue
			}
			if m == n {
				times[n]++
			}
			if m > n {
				break
			}
		}
	}
	sum2 := 0
	for i, t := range times {
		sum2 += i * t
	}

	println("Part 2: ", sum2)
}
