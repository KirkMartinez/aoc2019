package main

import (
	log "github.com/sirupsen/logrus"
	"strconv"
)

const day int = 9

func prelimTests() {
	type Test struct {
		input  string
		output []int
	}

	tests := [...]Test{
		Test{"109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99", []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}},
		Test{"1102,34915192,34915192,7,4,7,99,0", []int{1219070632396864}},
		Test{"104,1125899906842624,99", []int{1125899906842624}}}
	for id, test := range tests {
		result := runPart1(test.input)
		if !sliceMatch(result, test.output) {
			log.Printf("Expected %d, got %d", test.output, result)
			log.Fatal("Failed test #" + strconv.Itoa(id))
		}
	}
	log.Println("Prelim tests passed.")
}

func sliceMatch(s1, s2 []int) (match bool) {
	if len(s1) != len(s2) {
		return false
	}
	match = true
	for i := range s1 {
		match = match && s1[i] == s2[i]
	}
	return match
}

func runPart1(input string) []int {
	computer := initComputer(input, []int{1})
	computer = run(computer)
	return computer.outputs
}

func runPart2(input string) []int {
	computer := initComputer(input, []int{2})
	computer = run(computer)
	return computer.outputs
}

func part1() {
	input := getData(day)
	result := runPart1(input[0])
	log.Printf("Day 9, part 1 solution: %#v", result)
}

func part2() {
	input := getData(day)
	result := runPart2(input[0])
	log.Printf("Day 9, part 2 solution: %#v", result)
}

func main() {
	//log.SetLevel(log.InfoLevel)
	//log.SetLevel(log.DebugLevel)

	prelimTests()
	part1()
	part2()
}
