package main

import (
	"log"
	"strconv"
	"strings"
)

func prelimTests() {
	type Test struct {
		input  string
		output bool
	}

	tests := [3]Test{
		Test{"111111", true},
		Test{"223450", false},
		Test{"123789", false}}
	for id, test := range tests {
		result := meetsCriteria(test.input)
		if result != test.output {
			log.Println("Failed test #" + strconv.Itoa(id))
			log.Printf("Expected %t got %#v", test.output, result)
			log.Fatal("Stopping.")
		}
	}
	log.Println("Part 1 prelim tests passed.")
}

func meetsCriteria(val string) (result bool) {
	result = len(val) == 6
	result = result && dupDigits(val)
	result = result && nondecreasingDigits(val)
	return result
}

func meetsPart2Criteria(val string) (result bool) {
	result = len(val) == 6
	result = result && dupDigitsPart2(val)
	result = result && nondecreasingDigits(val)
	return result
}

func dupDigits(val string) (dups bool) {
	for i := 0; i < len(val)-1; i++ {
		if val[i] == val[i+1] {
			dups = true
			break
		}
	}
	return dups
}

func dupDigitsPart2(val string) (dups bool) {
	for i := 0; i < len(val)-1; i++ {
		if val[i] == val[i+1] {
			if i+1 == len(val)-1 && val[i-1] != val[i] ||
				i == 0 && val[i+2] != val[i] ||
				i > 0 && val[i-1] != val[i] && val[i+2] != val[i] {
				dups = true
				break
			}
		}
	}
	return dups
}

func nondecreasingDigits(val string) (result bool) {
	result = true
	for i := 0; i < len(val)-1; i++ {
		if val[i] > val[i+1] {
			result = false
			break
		}
	}
	return result
}

func runPart1(input string) (result int) {
	sf := strings.Split(input, "-")
	start, finish := sf[0], sf[1]

	startInt, err := strconv.Atoi(start)
	if err != nil {
		log.Fatal(err)
	}
	finishInt, err := strconv.Atoi(finish)
	if err != nil {
		log.Fatal(err)
	}

	for i := startInt; i <= finishInt; i++ {
		iStr := strconv.Itoa(i)
		if meetsCriteria(iStr) {
			result++
		}
	}

	return result
}

func runPart2(input string) (result int) {
	sf := strings.Split(input, "-")
	start, finish := sf[0], sf[1]

	startInt, err := strconv.Atoi(start)
	if err != nil {
		log.Fatal(err)
	}
	finishInt, err := strconv.Atoi(finish)
	if err != nil {
		log.Fatal(err)
	}

	for i := startInt; i <= finishInt; i++ {
		iStr := strconv.Itoa(i)
		if meetsPart2Criteria(iStr) {
			result++
		}
	}

	return result
}

func part1() {
	input := getData(4)
	result := runPart1(input[0])
	log.Printf("Day 4, part 1 solution: %d", result)
}

func part2() {
	input := getData(4)
	result := runPart2(input[0])
	log.Printf("Day 4, part 2 solution: %d", result)
}

func main() {
	prelimTests()
	part1()
	part2()
}
