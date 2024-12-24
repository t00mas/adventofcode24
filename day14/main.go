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

func yoloAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func main() {
	debug = flag.Bool("debug", false, "Debug")
	flag.Parse()

	DIRECTIONS := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	lines := strings.Split(string(input), "\n")
	robots := make([][4]int, len(lines))
	for i, line := range lines {
		pv := strings.Split(line, " ")
		pXY := strings.Split(strings.Split(pv[0], "=")[1], ",")
		vXY := strings.Split(strings.Split(pv[1], "=")[1], ",")
		robots[i] = [4]int{yoloAtoi(pXY[0]), yoloAtoi(pXY[1]), yoloAtoi(vXY[0]), yoloAtoi(vXY[1])}
	}

	X := 101
	Y := 103
	G := make([][]string, Y)
	for i := range Y {
		G[i] = make([]string, X)
	}

	p1, p2 := 0, 0
	for t := 1; t <= 100000; t++ {
		// Reset the grid
		for i := range Y {
			for j := range X {
				G[i][j] = ""
			}
		}
		q1, q2, q3, q4, mx, my := 0, 0, 0, 0, X/2, Y/2
		for i := range robots {
			px, py, vx, vy := robots[i][0], robots[i][1], robots[i][2], robots[i][3]
			px += vx
			py += vy
			px = (px + X) % X
			py = (py + Y) % Y
			robots[i] = [4]int{px, py, vx, vy}
			G[py][px] = "#"

			if t == 100 {
				mx = X / 2
				my = Y / 2
				if px < mx && py < my {
					q1++
				}
				if px > mx && py < my {
					q2++
				}
				if px < mx && py > my {
					q3++
				}
				if px > mx && py > my {
					q4++
				}
			}
		}

		if t == 100 {
			p1 = q1 * q2 * q3 * q4
		}

		components := 0
		seen := make(map[[2]int]bool)
		for x := 0; x < X; x++ {
			for y := 0; y < Y; y++ {
				if G[y][x] == "#" && !seen[[2]int{x, y}] {
					sx, sy := x, y
					components += 1
					queue := [][2]int{{sx, sy}}
					for len(queue) > 0 {
						x2, y2 := queue[0][0], queue[0][1]
						queue = queue[1:]
						if seen[[2]int{x2, y2}] {
							continue
						}
						seen[[2]int{x2, y2}] = true
						for _, dxy := range DIRECTIONS {
							xx, yy := x2+dxy[0], y2+dxy[1]
							if xx >= 0 && xx < X && yy >= 0 && yy < Y && G[yy][xx] == "#" {
								queue = append(queue, [2]int{xx, yy})
							}
						}
					}
				}
			}
		}

		if components <= 200 {
			p2 = t
			break
		}
	}

	fmt.Println("Part 1: ", p1)

	fmt.Println("Part 2: ", p2)
}
