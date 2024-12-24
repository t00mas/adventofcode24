package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	p1, p2 := 0, 0

	// One input at a time
	for i := 0; i < len(lines); i += 4 {
		parseButtonAValues := regexp.MustCompile(`X\+(-?\d+), Y\+(-?\d+)`).FindStringSubmatch(lines[i])
		AX, _ := strconv.Atoi(parseButtonAValues[1])
		AY, _ := strconv.Atoi(parseButtonAValues[2])

		parseButtonBValues := regexp.MustCompile(`X\+(-?\d+), Y\+(-?\d+)`).FindStringSubmatch(lines[i+1])
		BX, _ := strconv.Atoi(parseButtonBValues[1])
		BY, _ := strconv.Atoi(parseButtonBValues[2])

		parsePrizeValues := regexp.MustCompile(`X\=(-?\d+), Y\=(-?\d+)`).FindStringSubmatch(lines[i+2])
		PX, _ := strconv.Atoi(parsePrizeValues[1])
		PY, _ := strconv.Atoi(parsePrizeValues[2])

		PXp2 := PX + 10000000000000
		PYp2 := PY + 10000000000000

		countA := float64(PX*BY-BX*PY) / float64(BY*AX-BX*AY)
		countB := float64(PY*AX-AY*PX) / float64(BY*AX-BX*AY)

		countAp2 := float64(PXp2*BY-BX*PYp2) / float64(BY*AX-BX*AY)
		countBp2 := float64(PYp2*AX-AY*PXp2) / float64(BY*AX-BX*AY)

		if countA >= 0 && countB >= 0 && countA <= 100 && countB <= 100 && countA == float64(int(countA)) && countB == float64(int(countB)) {
			p1 += int(countA)*3 + int(countB)
		}
		if countAp2 >= 0 && countBp2 >= 0 && countAp2 == float64(int(countAp2)) && countBp2 == float64(int(countBp2)) {
			p2 += int(countAp2)*3 + int(countBp2)
		}
	}

	fmt.Printf("answer part 1: %d\n", p1)
	fmt.Printf("answer part 2: %d\n", p2)
}
