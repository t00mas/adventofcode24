package main

import (
	_ "embed"
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
	input := strings.Split(string(input), "\n")
	numRows := len(input)
	numCols := len(input[0])

	// find the guard
	var sr, sc int
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			if input[r][c] == '^' {
				sr, sc = r, c
			}
		}
	}

	p1 := 0
	p2 := 0

	for o_r := 0; o_r < numRows; o_r++ {
		for o_c := 0; o_c < numCols; o_c++ {
			r, c := sr, sc
			d := 0 // 0=up, 1=right, 2=down, 3=left
			seen := make(map[[3]int]bool)
			seenRC := make(map[[2]int]bool)
			for {
				if seen[[3]int{r, c, d}] {
					p2++
					break
				}
				seen[[3]int{r, c, d}] = true
				seenRC[[2]int{r, c}] = true
				dr, dc := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}[d][0], [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}[d][1]
				rr := r + dr
				cc := c + dc
				if !(0 <= rr && rr < numRows && 0 <= cc && cc < numCols) {
					if input[o_r][o_c] == '#' {
						p1 = len(seenRC)
					}
					break
				}
				if input[rr][cc] == '#' || (rr == o_r && cc == o_c) {
					d = (d + 1) % 4
				} else {
					r = rr
					c = cc
				}
			}
		}
	}

	fmt.Println("Part 1: ", p1)

	fmt.Println("Part 2: ", p2)
}
