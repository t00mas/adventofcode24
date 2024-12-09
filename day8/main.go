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
	debug = flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	grid := strings.Split(string(input), "\n")
	numRows := len(grid)
	numCols := len(grid[0])

	positions := parseGrid(grid, numRows, numCols)

	part1Matches := findMatches(positions, numRows, numCols, true)
	part2Matches := findMatches(positions, numRows, numCols, false)

	fmt.Printf("Part 1: %d matches\n", len(part1Matches))
	fmt.Printf("Part 2: %d matches\n", len(part2Matches))
}

func parseGrid(grid []string, numRows, numCols int) map[rune][][2]int {
	positions := make(map[rune][][2]int)
	for ridx := 0; ridx < numRows; ridx++ {
		for cidx := 0; cidx < numCols; cidx++ {
			if grid[ridx][cidx] != '.' {
				positions[rune(grid[ridx][cidx])] = append(positions[rune(grid[ridx][cidx])], [2]int{ridx, cidx})
			}
		}
	}
	return positions
}

func findMatches(positions map[rune][][2]int, numRows, numCols int, part1 bool) map[[2]int]struct{} {
	matches := make(map[[2]int]struct{})
	for _, coords := range positions {
		for i := 0; i < len(coords); i++ {
			for j := 0; j < len(coords); j++ {
				if i == j {
					continue
				}
				for ridx := 0; ridx < numRows; ridx++ {
					for cidx := 0; cidx < numCols; cidx++ {
						if isValidMatch(ridx, cidx, coords[i], coords[j], part1) {
							matches[[2]int{ridx, cidx}] = struct{}{}
						}
					}
				}
			}
		}
	}
	return matches
}

func isValidMatch(ridx, cidx int, coord1, coord2 [2]int, part1 bool) bool {
	ridx1, cidx1 := coord1[0], coord1[1]
	ridx2, cidx2 := coord2[0], coord2[1]

	distridx1 := ridx - ridx1
	distridx2 := ridx - ridx2
	distcidx1 := cidx - cidx1
	distcidx2 := cidx - cidx2

	// Check if the points are collinear
	if distridx1*distcidx2 != distcidx1*distridx2 {
		return false
	}

	d1 := abs(distridx1) + abs(distcidx1)
	d2 := abs(distridx2) + abs(distcidx2)

	if part1 {
		return d1 == 2*d2 || d1*2 == d2
	}
	return true
}

func abs(x int) int {
	return (x ^ (x >> 31)) - (x >> 31)
}
