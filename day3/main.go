package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input_txt []byte

var debug *bool

func p(a ...interface{}) {
	if *debug {
		fmt.Println(a...)
	}
}

func main() {
	debug = flag.Bool("debug", false, "Debug")
	flag.Parse()

	input := string(input_txt)

	// Part 1
	// Compile the regular expression to find all "mul(a,b)" patterns
	// and find all matches of the pattern in the input string
	reMul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	muls := reMul.FindAllStringSubmatch(input, -1)
	sum1 := 0
	for _, match := range muls {
		// Convert the captured groups (a and b) from strings to integers
		// and add the product of a and b to the sum
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		sum1 += a * b
	}
	fmt.Println("Part 1: ", sum1)

	// Part 2
	// Find all "do()" and "don't()" in the input string, keep the indexes
	reDo := regexp.MustCompile(`do\(\)`)
	doIndexes := reDo.FindAllStringIndex(input, -1)
	reDont := regexp.MustCompile(`don't\(\)`)
	dontIndexes := reDont.FindAllStringIndex(input, -1)
	// Find all the "mul(a,b)" indexes
	mulsIdx := reMul.FindAllStringIndex(input, -1)
	sum2 := 0
	for idx, match := range muls {
		mulIdx := mulsIdx[idx][0]
		p("mulIdx: ", mulIdx)
		doIdx := findLastIndexBefore(doIndexes, mulIdx)
		dontIdx := findLastIndexBefore(dontIndexes, mulIdx)
		// If the "don't()" index is before the "do()" index, add the product of a and b to the sum
		// because the "do()" is last before the "mul(a,b)"
		if dontIdx <= doIdx {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			sum2 += a * b
		}
	}
	fmt.Println("Part 2: ", sum2)
}

// findLastIndexBefore finds the last index before the limit
func findLastIndexBefore(indexes [][]int, limit int) int {
	lastIdx := 0
	for _, idx := range indexes {
		if idx[0] < limit {
			lastIdx = idx[0]
		} else {
			break
		}
	}
	return lastIdx
}
