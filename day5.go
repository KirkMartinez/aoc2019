package main

import (
	"log"
	"strconv"
)

const day int = 5

func prelimTests() {
	type Test struct {
		begState string
		input    int
		endState string
		output   int
	}

	tests := [...]Test{
		// Day 2 tests: addition and multiplication with indirect params; termination
		Test{"1,0,0,0,99", 0, "2,0,0,0,99", 0},
		Test{"2,3,0,3,99", 0, "2,3,0,6,99", 0},
		Test{"2,4,4,5,99,0", 0, "2,4,4,5,99,9801", 0},
		Test{"1,1,1,4,99,5,6,0,99", 0, "30,1,1,4,2,5,6,0,99", 0},
		// Day 5 tests
		Test{"3,0,4,0,99", 42, "42,0,4,0,99", 0},
		Test{"1002,4,3,4,33", 0, "1002,4,3,4,99", 0},
		Test{"1101,100,-1,4,0", 0, "1101,100,-1,4,99", 0},
		// Day 5 tests
		Test{"3,9,8,9,10,9,4,9,99,-1,8", 42, "", 0},
		Test{"3,9,8,9,10,9,4,9,99,-1,8", 8, "", 1},
		Test{"3,9,7,9,10,9,4,9,99,-1,8", 8, "", 0},
		Test{"3,9,7,9,10,9,4,9,99,-1,8", 7, "", 1},
		Test{"3,3,1108,-1,8,3,4,3,99", 0, "", 0},
		Test{"3,3,1108,-1,8,3,4,3,99", 8, "", 1},
		Test{"3,3,1107,-1,8,3,4,3,99", 7, "", 1},
		Test{"3,3,1107,-1,8,3,4,3,99", 8, "", 0},
		Test{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 0, "", 0},
		Test{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 2, "", 1},
		Test{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", -2, "", 1},
		Test{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 0, "", 0},
		Test{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 2, "", 1},
		Test{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", -2, "", 1},
		Test{"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", 7, "", 999},
		Test{"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", 8, "", 1000},
		Test{"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", 9, "", 1001}}
	for id, test := range tests {
		output, terminationState := runPart1(test.begState, test.input)
		if test.endState == "" {
			// We don't care about state, only output value
			if output != test.output {
				log.Printf("Expected %d, got %d", test.output, output)
				log.Fatal("Failed test #" + strconv.Itoa(id))
			}
		} else {
			if terminationState != test.endState {
				log.Printf("Expected end state " + test.endState + " got " + terminationState)
				log.Fatal("Failed test #" + strconv.Itoa(id))
			}
		}
	}
	log.Println("Day 5 prelim tests passed.")
}

func runPart1(begState string, input int) (output int, endState string) {
	computer := initComputer(begState, []int{input})
	computer = run(computer)
	output, endState = snapshotComputer(computer)
	return
}

func part1() {
	data := getData(day)
	out, _ := runPart1(data[0], 1) // 1 is input from Day 5, Part 1
	log.Printf("Day 5, Part 1 output: %d", out)
}

func part2() {
	data := getData(day)
	out, _ := runPart1(data[0], 5) // 5 is input from Day 5, Part 2
	log.Printf("Day 5, Part 2 output: %d", out)
}

func main() {
	prelimTests()
	part1()
	part2()
}
