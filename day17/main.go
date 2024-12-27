package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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

//go:embed input.txt
var input string

func yoloAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type State struct {
	A       uint64
	B       uint64
	C       uint64
	program []byte
	pointer uint64
	output  []byte
}

func (s *State) Init() {
	lines := strings.Split(input, "\n")
	fmt.Sscanf(lines[0], "Register A: %d", &s.A)
	fmt.Sscanf(lines[1], "Register B: %d", &s.B)
	fmt.Sscanf(lines[2], "Register C: %d", &s.C)

	var program string
	fmt.Sscanf(lines[4], "Program: %s", &program)
	for _, c := range program {
		if c == ',' {
			continue
		}
		s.program = append(s.program, byte(c)-'0')
	}
	s.pointer = 0
	s.output = []byte{}
}

// Combo operands 0 through 3 represent literal values 0 through 3.
// Combo operand 4 represents the value of register A.
// Combo operand 5 represents the value of register B.
// Combo operand 6 represents the value of register C.
// Combo operand 7 is reserved and will not appear in valid programs.
func (s *State) combooperand(opcode uint64) uint64 {
	if opcode > 3 {
		switch opcode {
		case 4:
			opcode = s.A
		case 5:
			opcode = s.B
		case 6:
			opcode = s.C
		}
	}
	return opcode
}

// The adv instruction (opcode 0) performs division. The numerator is the value in the A register.
// The denominator is found by raising 2 to the power of the instruction's combo operand.
// (So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.)
// The result of the division operation is truncated to an integer and then written to the A register.
func (s *State) adv(operand uint64) {
	operand = s.combooperand(operand)
	s.A /= uint64(math.Pow(2, float64(operand)))
}

// The bxl instruction (opcode 1) calculates the bitwise XOR of register B and the instruction's literal operand,
// then stores the result in register B.
func (s *State) bxl(operand uint64) {
	s.B = s.B ^ operand
}

// The bst instruction (opcode 2) calculates the value of its combo operand modulo 8 (thereby keeping only its lowest 3 bits),
// then writes that value to the B register.
func (s *State) bst(operand uint64) {
	s.B = s.combooperand(operand) % 8
}

// The jnz instruction (opcode 3) does nothing if the A register is 0.
// However, if the A register is not zero, it jumps by setting the instruction pointer to the value of its literal operand;
// if this instruction jumps, the instruction pointer is not increased by 2 after this instruction.
func (s *State) jnz(operand uint64) {
	if s.A != 0 {
		s.pointer = operand
	} else {
		s.pointer += 2
	}
}

// The bxc instruction (opcode 4) calculates the bitwise XOR of register B and register C,
// then stores the result in register B.
// (For legacy reasons, this instruction reads an operand but ignores it.)
func (s *State) bxc(_ uint64) {
	s.B = s.B ^ s.C
}

// The out instruction (opcode 5) calculates the value of its combo operand modulo 8,
// then outputs that value. (If a program outputs multiple values, they are separated by commas.)
func (s *State) out(operand uint64) {
	operand = s.combooperand(operand) % 8
	s.output = append(s.output, byte(operand))
}

// The bdv instruction (opcode 6) works exactly like the adv instruction except that the result is stored in the B register.
// (The numerator is still read from the A register.)
func (s *State) bdv(operand uint64) {
	operand = s.combooperand(operand)
	s.B = s.A / uint64(math.Pow(2, float64(operand)))
}

// The cdv instruction (opcode 7) works exactly like the adv instruction except that the result is stored in the C register.
// (The numerator is still read from the A register.)
func (s *State) cdv(operand uint64) {
	operand = s.combooperand(operand)
	s.C = s.A / uint64(math.Pow(2, float64(operand)))
}

func (s *State) Run() {
	for s.pointer+1 < uint64(len(s.program)) {
		o := uint64(s.program[s.pointer+1])
		switch s.program[s.pointer] {
		case 0:
			s.adv(o)
		case 1:
			s.bxl(o)
		case 2:
			s.bst(o)
		case 3:
			s.jnz(o)
			continue
		case 4:
			s.bxc(o)
		case 5:
			s.out(o)
		case 6:
			s.bdv(o)
		case 7:
			s.cdv(o)
		}
		s.pointer += 2
	}
}

// FindA finds the smallest A that will produce the same output as the program
func FindA(A uint64, i int) uint64 {
	s := State{}
	s.Init()
	s.A = A
	s.Run()

	if slices.Equal(s.output, s.program) {
		return A
	}
	if i == 0 || slices.Equal(s.output, s.program[len(s.program)-i:]) {
		for ni := range 8 {
			if na := FindA(8*A+uint64(ni), i+1); na > 0 {
				return na
			}
		}
	}

	return 0
}

func main() {
	debug = flag.Bool("debug", false, "Debug")
	flag.Parse()

	state := State{}
	state.Init()
	state.Run()
	fmt.Println("Part 1:", state.output)

	minA := FindA(0, 0)
	fmt.Println("Part 2:", minA)
}
