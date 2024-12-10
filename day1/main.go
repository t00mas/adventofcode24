package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

func main() {
	n1s, n2s := parseInput(string(input))

	slices.Sort(n1s)
	slices.Sort(n2s)

	sum1 := calculateAbsoluteDifferenceSum(n1s, n2s)
	fmt.Println("Part 1:", sum1)

	sum2 := calculateSumOfEqualProductsV2(n1s, n2s)
	fmt.Println("Part 2:", sum2)
}

// parseInput parses the input string and returns two slices of integers, one for each column of the input.
func parseInput(input string) ([]int, []int) {
	n1s := make([]int, 0)
	n2s := make([]int, 0)

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		ns := strings.Fields(line)
		if len(ns) != 2 {
			continue
		}

		n1, err := strconv.Atoi(ns[0])
		if err != nil {
			panic(err)
		}
		n1s = append(n1s, n1)

		n2, err := strconv.Atoi(ns[1])
		if err != nil {
			panic(err)
		}
		n2s = append(n2s, n2)
	}

	return n1s, n2s
}

// calculateAbsoluteDifferenceSum calculates the sum of the absolute differences between the elements of n1s and n2s.
func calculateAbsoluteDifferenceSum(n1s, n2s []int) int {
	totalDifference := 0
	for i := 0; i < len(n1s); i++ {
		totalDifference += abs(n1s[i] - n2s[i])
	}
	return totalDifference
}

// calculateSumOfEqualProductsV1 calculates the sum of the products of the elements of n1s and n2s that are equal.
// Uses a nested loop to compare each element of n1s with each element of n2s.
func calculateSumOfEqualProductsV1(n1s, n2s []int) int {
	equalCounts := make(map[int]int)
	for _, n1 := range n1s {
		for _, n2 := range n2s {
			if n2 < n1 {
				continue
			}
			if n2 == n1 {
				equalCounts[n1]++
			}
			if n2 > n1 {
				break
			}
		}
	}

	sum := 0
	for value, count := range equalCounts {
		sum += value * count
	}
	return sum
}

// calculateSumOfEqualProductsV2 calculates the sum of the products of the elements of n1s and n2s that are equal.
// Uses a map to store the counts of each element in n2s, then iterates over n1s and adds the product of the element
func calculateSumOfEqualProductsV2(n1s, n2s []int) int {
	equalCounts := make(map[int]int)
	for _, n2 := range n2s {
		equalCounts[n2]++
	}

	sum := 0
	for _, n1 := range n1s {
		if count, exists := equalCounts[n1]; exists {
			sum += n1 * count
		}
	}
	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
