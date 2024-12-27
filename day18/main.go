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
var input string

const N = 71

var DIRS = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // up right down left

func xy(line string) (int, int) {
	xy := strings.Split(line, ",")
	x, _ := strconv.Atoi(xy[0])
	y, _ := strconv.Atoi(xy[1])
	return x, y
}

func withinBounds(x, y int) bool {
	return 0 <= y && y < N && 0 <= x && x < N
}

func atExit(x, y int) bool {
	return x == N-1 && y == N-1
}

func main() {
	debug = flag.Bool("debug", false, "Debug")
	flag.Parse()

	G := make([][]byte, N)
	for i := range G {
		G[i] = make([]byte, N)
		for j := range G[i] {
			G[i][j] = '.'
		}
	}

	for i, line := range strings.Split(input, "\n") {
		x, y := xy(line)
		if withinBounds(x, y) {
			G[y][x] = '#'
		}

		QUEUE := [][3]int{{0, 0, 0}}
		SEEN := make(map[[2]int]bool)
		reachedExit := false
		for len(QUEUE) > 0 {
			d, row, col := QUEUE[0][0], QUEUE[0][1], QUEUE[0][2]
			QUEUE = QUEUE[1:]
			if atExit(row, col) {
				if i == 1023 {
					fmt.Println(d)
				}
				reachedExit = true
				break
			}
			if SEEN[[2]int{row, col}] {
				continue
			}
			SEEN[[2]int{row, col}] = true
			for _, direction := range DIRS {
				rr := row + direction[0]
				cc := col + direction[1]
				if withinBounds(rr, cc) && G[rr][cc] != '#' {
					QUEUE = append(QUEUE, [3]int{d + 1, rr, cc})
				}
			}
		}
		if !reachedExit {
			fmt.Printf("%d,%d\n", x, y)
			break
		}
	}
}
