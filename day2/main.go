package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input_txt []byte

func checkReport(report []string) bool {
	ld := "flat"
	for i := 1; i < len(report); i++ {
		a, _ := strconv.Atoi(report[i-1])
		b, _ := strconv.Atoi(report[i])
		diff := a - b
		var cd string
		if diff < 0 {
			cd = "inc"
		} else if diff > 0 {
			cd = "dec"
		}
		if ld == "flat" {
			ld = cd
		} else if ld != cd {
			return false
		}
		if diff == 0 {
			return false
		}
		if diff < 0 {
			diff = -diff
		}
		if diff > 3 {
			return false
		}
	}
	return true
}

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
	input := string(input_txt)
	lines := strings.Split(input, "\n")
	safe1 := len(lines)
	for _, line := range lines {
		ns := strings.Split(line, " ")
		if !checkReport(ns) {
			safe1--
			continue
		}
	}
	fmt.Println("Part 1: ", safe1)

	safe2 := 0
	for _, line := range lines {
		ns := strings.Split(line, " ")
		if checkReport(ns) {
			safe2++
			continue
		}
		if checkAllSubreports(ns) {
			safe2++
			continue
		}
	}
	fmt.Println("Part 2: ", safe2)
}
