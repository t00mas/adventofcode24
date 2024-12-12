package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inputfilled.txt
var input []byte

func findXMASaround(window [][]string) int {
	if window[3][3] != "X" {
		return 0
	}
	directions := [][][]int{
		{{2, 3}, {1, 3}, {0, 3}},
		{{2, 4}, {1, 5}, {0, 6}},
		{{3, 4}, {3, 5}, {3, 6}},
		{{4, 4}, {5, 5}, {6, 6}},
		{{4, 3}, {5, 3}, {6, 3}},
		{{4, 2}, {5, 1}, {6, 0}},
		{{3, 2}, {3, 1}, {3, 0}},
		{{2, 2}, {1, 1}, {0, 0}},
	}
	words := []string{"MAS", "MAS", "MAS", "MAS", "MAS", "MAS", "MAS", "MAS"}
	return findAroundWindow(window, directions, words)
}

func findMAScrossaround(window [][]string) int {
	if window[1][1] != "A" {
		return 0
	}
	directions := [][][]int{
		{{0, 0}, {2, 2}, {2, 0}, {0, 2}},
		{{0, 0}, {2, 2}, {2, 0}, {0, 2}},
		{{0, 0}, {2, 2}, {2, 0}, {0, 2}},
		{{0, 0}, {2, 2}, {2, 0}, {0, 2}},
	}
	words := []string{"MSMS", "MSSM", "SMMS", "SMSM"}
	return findAroundWindow(window, directions, words)
}

func findAroundWindow(window [][]string, directions [][][]int, words []string) int {
	count := len(directions)
	for i, directionCoords := range directions {
		word := words[i]
		for j, coord := range directionCoords {
			if window[coord[0]][coord[1]] != string(word[j]) {
				count--
				break
			}
		}
	}
	return count
}

func makeWindow(size int) [][]string {
	window := make([][]string, size)
	for i := range window {
		window[i] = make([]string, size)
	}
	return window
}

func main() {
	lines := strings.Split(string(input), "\n")

	count1 := 0
	for i := 4; i < len(lines)-3; i++ {
		for j := 3; j < len(lines[i])-3; j++ {
			window := makeWindow(7)
			for k := -3; k <= 3; k++ {
				for l := -3; l <= 3; l++ {
					window[k+3][l+3] = string(lines[i+k][j+l])
				}
			}
			count1 += findXMASaround(window)
		}
	}
	fmt.Println("Part 1: ", count1)

	count2 := 0
	for i := 4; i < len(lines)-4; i++ {
		for j := 4; j < len(lines[i])-4; j++ {
			window := makeWindow(3)
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					window[k+1][l+1] = string(lines[i+k][j+l])
				}
			}
			count2 += findMAScrossaround(window)
		}
	}
	fmt.Println("Part 2: ", count2)
}
