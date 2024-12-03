package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input_txt []byte

func main() {
	input := string(input_txt)

	reMul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	muls := reMul.FindAllStringSubmatch(input, -1)
	sum1 := 0
	for _, match := range muls {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		sum1 += a * b
	}

	fmt.Println("Part 1: ", sum1)

	reDo := regexp.MustCompile(`do\(\)`)
	doIndexes := reDo.FindAllStringIndex(input, -1)
	reDont := regexp.MustCompile(`don't\(\)`)
	dontIndexes := reDont.FindAllStringIndex(input, -1)
	mulsIdx := reMul.FindAllStringIndex(input, -1)

	sum2 := 0
	for idx, match := range muls {
		mulIdx := mulsIdx[idx][0]
		fmt.Println("mulIdx: ", mulIdx)
		doIdx := 0
		for _, dIdx := range doIndexes {
			if dIdx[0] < mulIdx {
				doIdx = dIdx[0]
				continue
			}
			break
		}
		dontIdx := 0
		for _, dIdx := range dontIndexes {
			if dIdx[0] < mulIdx {
				dontIdx = dIdx[0]
				continue
			}
			break
		}

		if dontIdx <= doIdx {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			sum2 += a * b
		}
	}

	fmt.Println("Part 2: ", sum2)
}
