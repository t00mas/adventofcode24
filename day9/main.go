package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
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

	p1 := solve(string(input), false)
	p2 := solve(string(input), true)

	fmt.Println("Part 1: ", p1)

	fmt.Println("Part 2: ", p2)
}

func solve(input string, isPart2 bool) int {
	files, spaces, final := parseInput(input, isPart2)
	final = rearrangeFiles(files, spaces, final)
	return checksum(final)
}

func parseInput(input string, isPart2 bool) ([][]int, [][]int, []*int) {
	var files [][]int
	var spaces [][]int
	var final []*int
	position := 0
	fileID := 0

	for i, char := range input {
		num, _ := strconv.Atoi(string(char))
		if i%2 == 0 {
			if isPart2 {
				files = append(files, []int{position, num, fileID})
			}
			for j := 0; j < num; j++ {
				id := fileID
				final = append(final, &id)
				if !isPart2 {
					files = append(files, []int{position, 1, fileID})
				}
				position++
			}
			fileID++
		} else {
			spaces = append(spaces, []int{position, num})
			for j := 0; j < num; j++ {
				final = append(final, nil)
				position++
			}
		}
	}
	return files, spaces, final
}

func rearrangeFiles(files [][]int, spaces [][]int, final []*int) []*int {
	for i := len(files) - 1; i >= 0; i-- {
		pos, size, fileID := files[i][0], files[i][1], files[i][2]
		for spaceIndex, space := range spaces {
			spacePos, spaceSize := space[0], space[1]
			if spacePos < pos && size <= spaceSize {
				for j := 0; j < size; j++ {
					if final[pos+j] == nil || *final[pos+j] != fileID {
						panic(fmt.Sprintf("final[pos+j]=%v", final[pos+j]))
					}
					final[pos+j] = nil
					final[spacePos+j] = &fileID
				}
				spaces[spaceIndex] = []int{spacePos + size, spaceSize - size}
				break
			}
		}
	}
	return final
}

func checksum(final []*int) int {
	result := 0
	for i, value := range final {
		if value != nil {
			result += i * *value
		}
	}
	return result
}
