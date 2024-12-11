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

func main() {
	debug = flag.Bool("debug", false, "Debug")
	flag.Parse()

	input := strings.Fields(string(input))
	inputRow := []int{}
	for _, v := range input {
		vInt, _ := strconv.Atoi(v)
		inputRow = append(inputRow, vInt)
	}
	p("Initial row")
	p(inputRow)

	row1 := []int{}
	row1 = append(row1, inputRow...)
	for i := 0; i < 25; i++ {
		p("Iteration", i)
		row1 = evolveStoneRowNaive(row1)
		p(row1)
	}
	numStones1 := len(row1)
	fmt.Println("Part 1: ", numStones1)

	row2 := []int{}
	row2 = append(row2, inputRow...)
	numStones2 := 0
	for _, v := range row2 {
		numStones2 += countEvolveStoneTimes(v, 75)
	}
	fmt.Println("Part 2: ", numStones2)
}

func evolveStoneRowNaive(row []int) []int {
	newRow := []int{}
	for _, v := range row {
		newRow = append(newRow, evolveStoneNaive(v)...)
	}
	return newRow
}

func evolveStoneNaive(v int) []int {
	if v == 0 {
		return []int{1}
	}
	vStr := strconv.Itoa(v)
	vStrLen := len(vStr)
	if vStrLen%2 == 0 {
		p1, _ := strconv.Atoi(vStr[:vStrLen/2])
		p2, _ := strconv.Atoi(vStr[vStrLen/2:])
		return []int{p1, p2}
	}
	return []int{v * 2024}
}

var DP = make(map[[2]int]int)

// countEvolveStoneTimes returns the number of stones after t iterations
// don't care about values
func countEvolveStoneTimes(x, t int) int {
	if val, ok := DP[[2]int{x, t}]; ok {
		// if value is memoized, return it
		return val
	}
	var ret int
	if t == 0 {
		// no more blinks to evolve the stone
		ret = 1
	} else if x == 0 {
		// stone value is 0, evolve to 1 and continue
		ret = countEvolveStoneTimes(1, t-1)
	} else if len(strconv.Itoa(x))%2 == 0 {
		// stone value length is even, split in half and evolve each half
		dstr := strconv.Itoa(x)
		left, _ := strconv.Atoi(dstr[:len(dstr)/2])
		right, _ := strconv.Atoi(dstr[len(dstr)/2:])
		ret = countEvolveStoneTimes(left, t-1) + countEvolveStoneTimes(right, t-1)
	} else {
		// stone value length is odd, multiply by 2024 and evolve
		ret = countEvolveStoneTimes(x*2024, t-1)
	}
	// memoize the result
	DP[[2]int{x, t}] = ret
	return ret
}
