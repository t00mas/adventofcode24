package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

var debug *bool

func p(a ...interface{}) {
	if *debug {
		fmt.Println(a...)
	}
}

var (
	rows, cols int
	grid       [][]int
	dp         map[[2]int]int
)

func main() {
	debug = flag.Bool("debug", false, "Debug")
	flag.Parse()

	// build the grid from the input
	lines := strings.Split(string(input), "\n")
	grid = make([][]int, len(lines))
	for i, line := range lines {
		row := strings.Split(line, "")
		grid[i] = make([]int, len(row))
		for j, val := range row {
			grid[i][j], _ = strconv.Atoi(val)
		}
	}

	rows = len(grid)
	cols = len(grid[0])
	dp = make(map[[2]int]int)

	var part1, part2 int
	// iterate over the grid
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// if the cell has value 9, calculate the number of ways to reach cells with value 0 from that one
			if grid[r][c] == 9 {
				part1 += ways1(r, c)
				part2 += ways2(r, c)
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

// ways1 calculates the number of ways to reach cells with value 0 from the starting cell using BFS
func ways1(startRow, startCol int) int {
	// count contains the number of found cells with value 0
	count := 0
	// queue contains the coordinates of the cells to visit
	queue := [][2]int{{startRow, startCol}}
	// seen contains the coordinates of the cells that have been visited
	seen := make(map[[2]int]bool)

	// while there are cells to visit
	for len(queue) > 0 {
		// pop the first cell from the queue
		r, c := queue[0][0], queue[0][1]
		queue = queue[1:]
		// if the cell has already been visited, skip it
		if seen[[2]int{r, c}] {
			continue
		}
		// mark the cell as visited
		seen[[2]int{r, c}] = true
		// if the cell has value 0, increment the count of found cells
		if grid[r][c] == 0 {
			count++
		}
		// add the neighbors of the cell to the queue
		for _, d := range [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			// rr, cc are the coordinates of the neighbor cell
			rr, cc := r+d[0], c+d[1]
			// if the neighbor cell is within the grid bounds and has value grid[r][c]-1, add it to the queue
			// because it's a valid trail (decrement of 1 from current cell value)
			if isValid(rr, cc) && grid[rr][cc] == grid[r][c]-1 {
				// add the neighbor cell to the queue
				queue = append(queue, [2]int{rr, cc})
			}
		}
	}
	return count
}

// ways2 calculates the number of ways to reach cells with value 0 from the starting cell using DFS with memoization
func ways2(r, c int) int {
	// if the cell has value 0, there's only one way to reach it
	if grid[r][c] == 0 {
		return 1
	}
	// if the number of ways to reach the cell has already been calculated, return it
	if val, ok := dp[[2]int{r, c}]; ok {
		return val
	}

	count := 0
	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for _, d := range directions {
		// rr, cc are the coordinates of the neighbor cell
		rr, cc := r+d[0], c+d[1]
		// if the neighbor cell is within the grid bounds and has value grid[r][c]-1, add it to the queue
		// because it's a valid trail (decrement of 1 from current cell value)
		if isValid(rr, cc) && grid[rr][cc] == grid[r][c]-1 {
			// recursively calculate the number of ways to reach the neighbor cell
			count += ways2(rr, cc)
		}
	}

	// memoize the result
	dp[[2]int{r, c}] = count
	return count
}

// isValid checks if the cell (r, c) is within the grid bounds
func isValid(r, c int) bool {
	return 0 <= r && r < rows && 0 <= c && c < cols
}
