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

	G := strings.Split(string(input), "\n")
	R := len(G)
	C := len(G[0])

	seen := make(map[[2]int]bool)
	p1, p2 := 0, 0
	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if seen[[2]int{r, c}] {
				continue
			}

			area, perimeter := 0, 0

			QUEUE := [][2]int{{r, c}}
			PERIMETER := make(map[[2]int]map[[2]int]bool)
			for len(QUEUE) > 0 {
				r2, c2 := QUEUE[0][0], QUEUE[0][1]
				QUEUE = QUEUE[1:]
				if seen[[2]int{r2, c2}] {
					continue
				}
				seen[[2]int{r2, c2}] = true
				area++
				for _, dir := range directions {
					rr, cc := r2+dir[0], c2+dir[1]
					if 0 <= rr && rr < R && 0 <= cc && cc < C && G[rr][cc] == G[r2][c2] {
						QUEUE = append(QUEUE, [2]int{rr, cc})
					} else {
						perimeter++
						if PERIMETER[dir] == nil {
							PERIMETER[dir] = make(map[[2]int]bool)
						}
						PERIMETER[dir][[2]int{r2, c2}] = true
					}
				}
			}

			sides := calculateSides(PERIMETER)

			p1 += area * perimeter
			p2 += area * sides
		}
	}

	fmt.Println("Part 1: ", p1)
	fmt.Println("Part 2: ", p2)
}

// calculateSides calculates the number of sides of the perimeter
// it uses a map to keep track of the visited perimeter points
// and a queue to traverse the perimeter
func calculateSides(perimeter map[[2]int]map[[2]int]bool) int {
	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	sides := 0
	for _, vs := range perimeter {
		seen := make(map[[2]int]bool)
		// prpc is the perimeter point coordinates (row / column)
		for prpc := range vs {
			if seen[prpc] {
				continue
			}
			sides++
			queue := [][2]int{prpc}
			for len(queue) > 0 {
				r2, c2 := queue[0][0], queue[0][1]
				queue = queue[1:]
				if seen[[2]int{r2, c2}] {
					continue
				}
				seen[[2]int{r2, c2}] = true
				for _, direction := range directions {
					rr, cc := r2+direction[0], c2+direction[1]
					if vs[[2]int{rr, cc}] {
						queue = append(queue, [2]int{rr, cc})
					}
				}
			}
		}
	}
	return sides
}
