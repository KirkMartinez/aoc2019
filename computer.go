// Code related to the ship computer.
package main

import (
	"log"
	"strconv"
	"strings"
)

type IntComputer struct {
	inputs, outputs, state []int
	ip                     int // instruction pointer
	terminated             bool
}

type ParamMode int

const (
	POS ParamMode = iota
	DIRECT
)

// Converts a computer state string into a slice of integer opcodes and parameters
func initComputer(state string, ins []int) (computer IntComputer) {
	computer.inputs = ins

	computer.ip = 0

	inst := strings.Split(state, ",")
	computer.state = make([]int, len(inst))
	for id, item := range inst {
		opcode, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal("Can't convert instruction to int")
		}
		computer.state[id] = opcode
	}

	return computer
}

// Return current state of computer as a string
func snapshotComputer(computer IntComputer) (int, string) {
	strCode := make([]string, len(computer.state))
	for id, item := range computer.state {
		strCode[id] = strconv.Itoa(item)
	}

	var output int
	if len(computer.outputs) > 0 {
		output = computer.outputs[0]
	}
	return output, strings.Join(strCode, ",")
}

// Given a computer, run the program it contains
func run(computer IntComputer) IntComputer {
	//log.Printf("In: %d, Out: %d, Computer: %#v", input, output, computer)
	op := computer.state[computer.ip]
	if op == 99 {
		computer.terminated = true
		return computer
	}

	computer = processOp(op, computer)
	computer = run(computer)

	return computer
}

// Given a computer, execute the next instruction
func runStep(computer IntComputer) IntComputer {
	op := computer.state[computer.ip]
	if op == 99 {
		computer.terminated = true
		return computer
	}

	return processOp(op, computer)
}

func willTerminate(computer IntComputer) bool {
	return computer.state[computer.ip] == 99
}

func opAdd(pm [4]ParamMode, computer IntComputer) IntComputer {
	pc := computer.ip
	p1 := getParamValue(pm[1], pc+1, computer)
	p2 := getParamValue(pm[2], pc+2, computer)
	pdest := computer.state[pc+3]
	//log.Printf("opAdd %d(@%d) + %d(@%d) into %d\n", p1, pc+1, p2, pc+2, pc+3)
	computer.state[pdest] = p1 + p2
	computer.ip += 4
	return computer
}

func opMult(pm [4]ParamMode, computer IntComputer) IntComputer {
	pc := computer.ip
	p1 := getParamValue(pm[1], pc+1, computer)
	p2 := getParamValue(pm[2], pc+2, computer)
	pdest := computer.state[pc+3]
	//log.Printf("opMult %d(@%d) + %d(@%d) into %d\n", p1, pc+1, p2, pc+2, pc+3)
	computer.state[pdest] = p1 * p2
	computer.ip += 4
	return computer
}

func opReadInput(computer IntComputer) IntComputer {
	//log.Printf("opReadInput for %#v\n", computer)
	pc := computer.ip
	pdest := computer.state[pc+1]

	// Try to read input
	var input int
	if len(computer.inputs) > 0 {
		// Pop an input value
		n_inputs := len(computer.inputs)
		input, computer.inputs = computer.inputs[n_inputs-1], computer.inputs[:n_inputs-1]
		//log.Printf("Read input value %d\n", input)
	}

	computer.state[pdest] = input
	computer.ip += 2
	return computer
}

func opWriteOutput(pm [4]ParamMode, computer IntComputer) IntComputer {
	pc := computer.ip
	p1 := getParamValue(pm[1], pc+1, computer)
	//log.Printf("opWriteOutput: %#v\n", p1)
	computer.outputs = append(computer.outputs, p1)
	computer.ip += 2
	return computer
}

func opJumpIfTrue(pm [4]ParamMode, computer IntComputer) IntComputer {
	pc := computer.ip
	p1 := getParamValue(pm[1], pc+1, computer)
	p2 := getParamValue(pm[2], pc+2, computer)
	if p1 != 0 {
		computer.ip = p2
	} else {
		computer.ip += 3
	}

	return computer
}

func opJumpIfFalse(pm [4]ParamMode, computer IntComputer) IntComputer {
	pc := computer.ip
	p1 := getParamValue(pm[1], pc+1, computer)
	p2 := getParamValue(pm[2], pc+2, computer)
	if p1 == 0 {
		computer.ip = p2
	} else {
		computer.ip += 3
	}

	return computer
}

func opLessThan(pm [4]ParamMode, computer IntComputer) IntComputer {
	pc := computer.ip
	p1 := getParamValue(pm[1], pc+1, computer)
	p2 := getParamValue(pm[2], pc+2, computer)
	pdest := computer.state[pc+3]
	if p1 < p2 {
		computer.state[pdest] = 1
	} else {
		computer.state[pdest] = 0
	}
	computer.ip += 4

	return computer
}

func opEquals(pm [4]ParamMode, computer IntComputer) IntComputer {
	pc := computer.ip
	p1 := getParamValue(pm[1], pc+1, computer)
	p2 := getParamValue(pm[2], pc+2, computer)
	pdest := computer.state[pc+3]
	if p1 == p2 {
		computer.state[pdest] = 1
	} else {
		computer.state[pdest] = 0
	}
	computer.ip += 4

	return computer
}

func processOp(op int, computer IntComputer) IntComputer {
	opcode, paramModes := decodeOp(op)

	switch opcode {
	case 1:
		computer = opAdd(paramModes, computer)
	case 2:
		computer = opMult(paramModes, computer)
	case 3:
		computer = opReadInput(computer)
	case 4:
		computer = opWriteOutput(paramModes, computer)
	case 5:
		computer = opJumpIfTrue(paramModes, computer)
	case 6:
		computer = opJumpIfFalse(paramModes, computer)
	case 7:
		computer = opLessThan(paramModes, computer)
	case 8:
		computer = opEquals(paramModes, computer)
	default:
		log.Fatal("Unexpected opcode!")
	}

	return computer
}

func decodeOp(op int) (opcode int, paramModes [4]ParamMode) {
	opcode = op % 100

	paramModes[0] = 999 // Param 0 is not well-defined; only 1-3 are valid
	pval := (op - opcode) / 100
	for p := 1; p < 4; p++ {
		if pval == 0 {
			paramModes[p] = 0
		} else {
			digit := pval % 10
			paramModes[p] = ParamMode(digit)
			pval = (pval - digit) / 10
		}
	}
	return
}

func getParamValue(mode ParamMode, loc int, computer IntComputer) int {
	switch mode {
	case POS:
		return computer.state[computer.state[loc]]
	case DIRECT:
		return computer.state[loc]
	}
	log.Fatal("Unknown parameter mode ", mode)
	return 0
}
