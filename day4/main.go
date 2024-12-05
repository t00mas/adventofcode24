package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inputfilled.txt
var input_txt []byte

func findXMASaround(window [7][7]string) int {
	word := "XMAS"
	seqs := map[string][][]int{
		"N ": {{3, 3}, {2, 3}, {1, 3}, {0, 3}},
		"NE": {{3, 3}, {2, 4}, {1, 5}, {0, 6}},
		"E ": {{3, 3}, {3, 4}, {3, 5}, {3, 6}},
		"SE": {{3, 3}, {4, 4}, {5, 5}, {6, 6}},
		"S ": {{3, 3}, {4, 3}, {5, 3}, {6, 3}},
		"SW": {{3, 3}, {4, 2}, {5, 1}, {6, 0}},
		"W ": {{3, 3}, {3, 2}, {3, 1}, {3, 0}},
		"NW": {{3, 3}, {2, 2}, {1, 1}, {0, 0}},
	}
	count := len(seqs)
	for _, seq := range seqs {
		for i, coord := range seq {
			if window[coord[0]][coord[1]] != string(word[i]) {
				count--
				break
			}
		}
	}
	return count
}

func findMAScrossaround(window [3][3]string) int {
	if window[1][1] != "A" {
		return 0
	}
	seqs := map[string][][]int{
		// "h_MASMS": {{1, 0}, {1, 1}, {1, 2}, {0, 1}, {2, 1}},
		// "h_MASSM": {{1, 0}, {1, 1}, {1, 2}, {0, 1}, {2, 1}},
		// "h_SAMMS": {{1, 0}, {1, 1}, {1, 2}, {0, 1}, {2, 1}},
		// "h_SAMSM": {{1, 0}, {1, 1}, {1, 2}, {0, 1}, {2, 1}},
		"d_MASMS": {{0, 0}, {1, 1}, {2, 2}, {2, 0}, {0, 2}},
		"d_MASSM": {{0, 0}, {1, 1}, {2, 2}, {2, 0}, {0, 2}},
		"d_SAMMS": {{0, 0}, {1, 1}, {2, 2}, {2, 0}, {0, 2}},
		"d_SAMSM": {{0, 0}, {1, 1}, {2, 2}, {2, 0}, {0, 2}},
	}
	count := len(seqs)
	for k, seq := range seqs {
		word := k[2:]
		for i, coord := range seq {
			if window[coord[0]][coord[1]] != string(word[i]) {
				count--
				break
			}
		}
	}
	return count
}

func main() {
	input := string(input_txt)
	lines := strings.Split(input, "\n")
	count1 := 0
	for i := 4; i < len(lines)-3; i++ {
		for j := 3; j < len(lines[i])-3; j++ {
			window := [7][7]string{}
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
			window := [3][3]string{}
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
