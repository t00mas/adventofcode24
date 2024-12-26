package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strings"
)

var debug *bool

func p(a ...interface{}) {
	if *debug {
		fmt.Println(a...)
	}
}

//go:embed input.txt
var input string

type D struct{ r, c int } // direction
type P struct{ r, c int }

var (
	E  = D{c: 1}
	S  = D{r: 1}
	W  = D{c: -1}
	N  = D{r: -1}
	Ds = []D{E, S, W, N}
)

func (d D) TurnRight() D {
	switch d {
	case E:
		return S
	case S:
		return W
	case W:
		return N
	case N:
		return E
	default:
		return d
	}
}

func (d D) TurnLeft() D {
	switch d {
	case E:
		return N
	case N:
		return W
	case W:
		return S
	case S:
		return E
	default:
		return d
	}
}

func (p P) Move(d D) P {
	return P{r: p.r + d.r, c: p.c + d.c}
}

type St struct {
	p P
	d D
}

func (s St) Possible() (str8, left, right St) {
	str8 = St{p: s.p.Move(s.d), d: s.d}
	left = St{p: s.p, d: s.d.TurnLeft()}
	right = St{p: s.p, d: s.d.TurnRight()}
	return
}

type O struct {
	c int
	p []St
}

// AddP adds a point to the list of points with the lowest cost.
func (o *O) AddP(p St, c int) {
	if o.c > c {
		o.c = c
		o.p = []St{p}
	} else if o.c == c {
		o.p = append(o.p, p)
	}
}

type Solver struct {
	G        []string
	pq       map[int][]St
	cheapest int
	highest  int
	end      St
	V        map[St]int
	orig     map[St]*O
}

func (s *Solver) A(v, prev St, cost int) {
	if cost < s.cheapest {
		panic("cost < s.cheapest")
	}
	p := s.orig[v]
	if p == nil {
		p = &O{c: cost}
		s.orig[v] = p
	}
	p.AddP(prev, cost)
	if c, ok := s.V[v]; !ok || cost < c {
		s.V[v] = cost
		s.pq[cost] = append(s.pq[cost], v)
		if cost > s.highest {
			s.highest = cost
		}
	}
}

func (s *Solver) printG() {
	g := make([][]byte, len(s.G))
	for r, l := range s.G {
		g[r] = []byte(strings.Clone(l))
	}
	q := []St{s.end}
	var zero St
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		if v != zero {
			q = append(q, s.orig[v].p...)
		}
		var d byte
		switch v.d {
		case E:
			d = '>'
		case W:
			d = '<'
		case N:
			d = '^'
		case S:
			d = 'v'
		}
		g[v.p.r][v.p.c] = d
	}
	for _, l := range g {
		p(string(l))
	}
}

func (s *Solver) pop(cost int) St {
	v := s.pq[cost][0]
	s.pq[cost] = s.pq[cost][1:]
	return v
}

func (s *Solver) Get(p P) byte {
	return s.G[p.r][p.c]
}

func Path(G []string, start St) *Solver {
	s := &Solver{G: G, pq: map[int][]St{}, V: map[St]int{}, orig: map[St]*O{}}
	s.A(start, St{}, 0)
	for {
		for len(s.pq[s.cheapest]) == 0 {
			if s.cheapest > s.highest {
				log.Fatalf("Ran out of priority queue: %d > %d", s.cheapest, s.highest)
			}
			s.cheapest++
		}
		v := s.pop(s.cheapest)
		if s.Get(v.p) == 'E' {
			s.end = v
			return s
		}
		straight, left, right := v.Possible()
		if s.Get(straight.p) != '#' {
			s.A(straight, v, s.cheapest+1)
		}
		if s.Get(left.p) != '#' {
			s.A(left, v, s.cheapest+1000)
		}
		if s.Get(right.p) != '#' {
			s.A(right, v, s.cheapest+1000)
		}
	}
}

func P1(lines []string) string {
	start := St{p: P{r: len(lines) - 2, c: 1}, d: E}
	if lines[start.p.r][start.p.c] != 'S' {
		start = St{p: P{r: 1, c: len(lines[0]) - 2}, d: S}
	}
	s := Path(lines, start)
	s.printG()
	return fmt.Sprintf("%d", s.cheapest)
}

func P2(lines []string) string {
	start := St{p: P{r: len(lines) - 2, c: 1}, d: E}
	if lines[start.p.r][start.p.c] != 'S' {
		start = St{p: P{r: 1, c: len(lines[0]) - 2}, d: S}
	}
	s := Path(lines, start)
	seen := make(map[P]bool)
	q := []St{s.end}
	var zero St
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		if v != zero {
			seen[v.p] = true
			q = append(q, s.orig[v].p...)
		}
	}
	return fmt.Sprintf("%d", len(seen))
}

func main() {
	debug = flag.Bool("debug", false, "Debug")
	flag.Parse()

	lines := strings.Split(input, "\n")
	fmt.Println(P1(lines))
	fmt.Println(P2(lines))
}
