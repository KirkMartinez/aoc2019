package main

import (
	"log"
	"strconv"
)

const day int = PUT_THE_DAY_HERE

func prelimTests() {
	type Test struct {
		input  string
		output int
	}

	tests := [...]Test{
		Test{"", 42}}
	for id, test := range tests {
		if runPart1(test.input) != test.output {
			log.Println("Failed test #" + strconv.Itoa(id))
			log.Fatal("Expected " + test.output + " got " + runPart1(test.input))
		}
	}
	log.Println("Prelim tests passed.")
}

func runPart1(input string) (result string) {
	return
}

func part1() {
	input := getData(day)
	result := runPart1(input)
	log.Println(result)
}

func part2() {
}

func main() {
	prelimTests()
	part1()
	part2()
}
