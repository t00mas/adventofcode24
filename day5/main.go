package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var debug *bool

func p(a ...interface{}) {
	if *debug {
		fmt.Println(a...)
	}
}

//go:embed inputrules.txt
var rules []byte

//go:embed inputpages.txt
var pages []byte

func main() {
	debug = flag.Bool("debug", false, "Debug")
	flag.Parse()

	rules := strings.Split(string(rules), "\n")
	afters := make(map[string][]string)
	befores := make(map[string][]string)
	for _, rule := range rules {
		pair := strings.Split(rule, "|")
		afters[pair[0]] = append(afters[pair[0]], pair[1])
		befores[pair[1]] = append(befores[pair[1]], pair[0])
	}

	pages := strings.Split(string(pages), "\n")
	sumMiddles := 0
	disorderedPages := make([]string, 0)
PAGE:
	for _, page := range pages {
		p("\nChecking", page)
		pageNums := strings.Split(page, ",")
		bf := make([]string, 0)
		af := pageNums[1:]
		for i := 0; i < len(pageNums); i++ {
			pageNum := pageNums[i]
			p("	Current", pageNum)
			p("	Befores:", bf)
			p("	Afters:", af)

			p("	", len(af))
			for j := 0; j < len(af); j++ {
				p("		Checking", af[j], "in", pageNum, befores[pageNum])
				if slices.Contains(befores[pageNum], af[j]) {
					p("		Found", af[j], "in", pageNum, befores[pageNum])
					disorderedPages = append(disorderedPages, page)
					continue PAGE
				}
			}

			p("	", len(bf))
			for j := 0; j < len(bf); j++ {
				p("		Checking", bf[j], "in", pageNum, afters[pageNum])
				if slices.Contains(afters[pageNum], bf[j]) {
					p("		Found", bf[j], "in", pageNum, afters[pageNum])
					disorderedPages = append(disorderedPages, page)
					continue PAGE
				}
			}

			bf = append(bf, pageNum)
			if len(af) > 0 {
				af = af[1:]
			}
		}

		mid, _ := strconv.Atoi(pageNums[((len(pageNums) - 1) / 2)])
		sumMiddles += mid
	}

	fmt.Println("Part 1: ", sumMiddles)

	sortFn := func(i, j string) int {
		if slices.Contains(befores[i], j) || slices.Contains(afters[j], i) {
			return -1
		}
		if slices.Contains(afters[i], j) || slices.Contains(befores[j], i) {
			return 1
		}
		return 0
	}
	sumMiddlesDisordered := 0
	for i := 0; i < len(disorderedPages); i++ {
		pageNums := strings.Split(disorderedPages[i], ",")
		slices.SortFunc(pageNums, sortFn)
		mid, _ := strconv.Atoi(pageNums[((len(pageNums) - 1) / 2)])
		sumMiddlesDisordered += mid
	}

	fmt.Println("Part 2: ", sumMiddlesDisordered)
}
