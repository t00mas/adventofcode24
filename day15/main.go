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
var input string

func parseInput(input string) ([][]string, [][]string, string) {
	blocks := strings.Split(input, "\n\n")

	lines := strings.Split(blocks[0], "\n")
	grid := make([][]string, 0)
	for _, line := range lines {
		row := []string(strings.Split(line, ""))
		p(row)
		grid = append(grid, row)
	}

	instructions := blocks[1]
	p(instructions)

	grid2 := make([][]string, 0)
	for i := 0; i < len(grid); i++ {
		r := ""
		for j := 0; j < len(grid[i]); j++ {
			switch grid[i][j] {
			case "#":
				r = fmt.Sprintf("%s##", r)
				break
			case ".":
				r = fmt.Sprintf("%s..", r)
				break
			case "O":
				r = fmt.Sprintf("%s[]", r)
				break
			case "@":
				r = fmt.Sprintf("%s@.", r)
				break
			}
		}
		row := []string(strings.Split(r, ""))
		p(row)
		grid2 = append(grid2, row)
	}

	return grid, grid2, instructions
}

func solveP1(grid [][]string, instruction string) int {
	xy := [2]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "@" {
				xy = [2]int{i, j}
			}
		}
	}

	for _, i := range []byte(instruction) {
		xy = robotDoV1(i, xy, grid)
	}

	return acc(grid, "O")
}

func solveP2(grid [][]string, instruction string) int {
	xy := [2]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "@" {
				xy = [2]int{i, j}
			}
		}
	}

	for _, i := range []byte(instruction) {
		if strings.Contains("^v", string(i)) {
			grid2 := gridCopy(grid)
			newxy := robotDoV2(i, xy, xy, grid2)
			if newxy[0] != -1 {
				xy = newxy
				grid = grid2
			}
		} else {
			xy = robotDoV1(i, xy, grid)
		}
	}

	return acc(grid, "[")
}

// grioCopy deepcopies the grid
func gridCopy(grid [][]string) [][]string {
	newGrid := make([][]string, 0)
	for i := 0; i < len(grid); i++ {
		newGrid = append(newGrid, make([]string, len(grid[i])))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}

// robotDoV1 moves the robot in the grid
func robotDoV1(ins byte, xy [2]int, grid [][]string) [2]int {
	cxy := xy
	for {
		xxyy := xyWalk(cxy, ins, false)
		if grid[xxyy[0]][xxyy[1]] == "#" {
			return xy
		} else if grid[xxyy[0]][xxyy[1]] == "." {
			cxy = xxyy
			break
		}
		if cxy == xxyy {
			break
		}
		cxy = xxyy
	}

	for {
		nxy := xyWalk(cxy, ins, true)
		grid[cxy[0]][cxy[1]] = grid[nxy[0]][nxy[1]]
		if nxy == xy {
			grid[xy[0]][xy[1]] = "."
			grid[cxy[0]][cxy[1]] = "@"
			return cxy
		}
		cxy = nxy
	}
}

// robotDoV2 is a recursive function that moves the robot in the grid
func robotDoV2(ins byte, xy1 [2]int, xy2 [2]int, grid [][]string) [2]int {
	if xy1 == xy2 {
		nxy := xyWalk(xy1, ins, false)
		if grid[nxy[0]][nxy[1]] == "#" {
			return [2]int{-1, -1}
		}
		if grid[nxy[0]][nxy[1]] == "[" {
			newPos := robotDoV2(ins, nxy, [2]int{nxy[0], nxy[1] + 1}, grid)
			if newPos[0] == -1 {
				return [2]int{-1, -1}
			}
		}
		if grid[nxy[0]][nxy[1]] == "]" {
			newPos := robotDoV2(ins, [2]int{nxy[0], nxy[1] - 1}, nxy, grid)
			if newPos[0] == -1 {
				return [2]int{-1, -1}
			}
		}

		grid[nxy[0]][nxy[1]] = grid[xy1[0]][xy1[1]]
		grid[xy1[0]][xy1[1]] = "."
		return nxy
	}

	nxy1 := xyWalk(xy1, ins, false)
	nxy2 := xyWalk(xy2, ins, false)

	if grid[nxy1[0]][nxy1[1]] == "#" || grid[nxy2[0]][nxy2[1]] == "#" {
		return [2]int{-1, -1}
	}
	if grid[nxy1[0]][nxy1[1]] == "[" {
		newPos := robotDoV2(ins, nxy1, nxy2, grid)
		if newPos[0] == -1 {
			return [2]int{-1, -1}
		}
	}
	if grid[nxy1[0]][nxy1[1]] == "]" {
		newPos := robotDoV2(ins, [2]int{nxy1[0], nxy1[1] - 1}, nxy1, grid)
		if newPos[0] == -1 {
			return [2]int{-1, -1}
		}
	}
	if grid[nxy2[0]][nxy2[1]] == "[" {
		newPos := robotDoV2(ins, nxy2, [2]int{nxy2[0], nxy2[1] + 1}, grid)
		if newPos[0] == -1 {
			return [2]int{-1, -1}
		}
	}

	grid[nxy1[0]][nxy1[1]] = grid[xy1[0]][xy1[1]]
	grid[xy1[0]][xy1[1]] = "."

	grid[nxy2[0]][nxy2[1]] = grid[xy2[0]][xy2[1]]
	grid[xy2[0]][xy2[1]] = "."

	return xy1
}

func xyWalk(xy [2]int, dir byte, flip bool) [2]int {
	if flip {
		switch dir {
		case '<':
			dir = '>'
			break
		case '>':
			dir = '<'
			break
		case '^':
			dir = 'v'
			break
		case 'v':
			dir = '^'
			break
		}
	}

	x, y := xy[0], xy[1]
	switch dir {
	case '<':
		return [2]int{x, y - 1}
	case '>':
		return [2]int{x, y + 1}
	case '^':
		return [2]int{x - 1, y}
	case 'v':
		return [2]int{x + 1, y}
	}

	return xy
}

func acc(grid [][]string, char string) int {
	s := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == char {
				s += (100 * i) + j
			}
		}
	}
	return s
}

func main() {
	debug = flag.Bool("debug", false, "Debug")
	flag.Parse()

	g, g2, instructions := parseInput(input)
	fmt.Println("Part 1: ", solveP1(g, instructions))
	fmt.Println("Part 2: ", solveP2(g2, instructions))
}
