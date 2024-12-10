package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input []byte

// checkReport checks if the report is valid (within 3 steps and no flat steps, respecting the order)
func checkReport(report []string) bool {
	lastDirection := "flat"
	for i := 1; i < len(report); i++ {
		prev, _ := strconv.Atoi(report[i-1])
		curr, _ := strconv.Atoi(report[i])
		diff := prev - curr
		var currDirection string
		if diff < 0 {
			currDirection = "inc"
		} else if diff > 0 {
			currDirection = "dec"
		}

		if lastDirection == "flat" {
			lastDirection = currDirection
		} else if lastDirection != currDirection {
			return false
		}

		if diff == 0 || abs(diff) > 3 {
			return false
		}
	}
	return true
}

// checkAllSubreports checks if the report is still valid if removing any one element
func checkAllSubreports(report []string) bool {
	for i := 0; i < len(report); i++ {
		subReport := make([]string, 0)
		subReport = append(subReport, report[:i]...)
		subReport = append(subReport, report[i+1:]...)
		if checkReport(subReport) {
			return true
		}
	}
	return false
}

func main() {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	numSafeReports1 := 0
	numSafeReports2 := 0
	for _, line := range lines {
		ns := strings.Fields(line)
		if checkReport(ns) {
			numSafeReports1++
		}
		if checkReport(ns) || checkAllSubreports(ns) {
			numSafeReports2++
		}
	}
	fmt.Println("Part 1:", numSafeReports1)
	fmt.Println("Part 2:", numSafeReports2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
