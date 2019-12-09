package main

import (
	"log"
	"strconv"
	"strings"
)

func prelimTests() {
	type Test struct {
		input, output string
	}

	tests := [4]Test{
		Test{"1,0,0,0,99", "2,0,0,0,99"},
		Test{"2,3,0,3,99", "2,3,0,6,99"},
		Test{"2,4,4,5,99,0", "2,4,4,5,99,9801"},
		Test{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"}}
	for id, test := range tests {
		if runPart1(test.input, false, 0, 0) != test.output {
			log.Println("Failed test #" + strconv.Itoa(id))
			log.Fatal("Expected " + test.output + " got " + runPart1(test.input, false, 0, 0))
		}
	}
	log.Println("Day 2 prelim tests passed.")
}

func makeComputer(inst []string) []int {
	var comp []int = make([]int, len(inst))
	for id, item := range inst {
		opcode, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal("Can't convert instruction to int")
		}
		comp[id] = opcode
	}
	return comp
}

func snapshotComputer(comp []int) []string {
	var result []string = make([]string, len(comp))
	for id, item := range comp {
		result[id] = strconv.Itoa(item)
	}
	return result
}

func runPart1(input string, fix bool, noun int, verb int) string {
	instructions := strings.Split(input, ",")
	computer := makeComputer(instructions)
	if fix {
		computer[1] = noun
		computer[2] = verb
	}
	process(computer, 0)
	snapshot := snapshotComputer(computer)
	return strings.Join(snapshot, ",")
}

func process(computer []int, pc int) {
	op := computer[pc]
	if op == 99 {
		return
	}
	psrc1 := computer[pc+1]
	psrc2 := computer[pc+2]
	pdest := computer[pc+3]

	if op == 1 {
		computer[pdest] = computer[psrc1] + computer[psrc2]
	} else if op == 2 {
		computer[pdest] = computer[psrc1] * computer[psrc2]
	} else {
		log.Fatal("Unexpected opcode!")
	}

	process(computer, pc+4)
}

func part1() {
	input := getData(2)
	result := runPart1(input[0], true, 12, 2)
	ops := strings.Split(result, ",")
	log.Printf("Day 2, part 1 solution: %s\n", ops[0])
}

func part2() {
	input := getData(2)
	initialState := input[:]

	var noun, verb int
	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			result := runPart1(input[0], true, noun, verb)
			state := strings.Split(result, ",")
			if state[0] == "19690720" {
				log.Printf("Day 2, part 2 solution: noun=%d, verb=%d\n", noun, verb)
				break
			}
			input = initialState[:]
		}
	}
}

func main() {
	prelimTests()
	part1()
	part2()
}
