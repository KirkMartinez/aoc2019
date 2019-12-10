// Code related to the ship computer.
//
// Instruction set:
//
// Opcode 1 adds together numbers read from two positions and stores the result in a third position.
// Opcode 2 works exactly like opcode 1, except it multiplies the two inputs instead of adding them.
// Opcode 3 takes a single integer as input and saves it to the position given by its only parameter.
// Opcode 4 outputs the value of its only parameter.
// Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter.
// Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the second parameter.
// Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
// Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
// Opcode 9 adjusts the relative base by the value of its only parameter. The relative base increases (or decreases, if the value is negative) by the value of the parameter.
//
// Parameter modes:
// Mode 0, position mode, causes the parameter to be interpreted as a position - if the parameter is 50, its value is the value stored at address 50 in memory.
// Mode 1, immediate mode, causes a parameter to be interpreted as a value - if the parameter is 50, its value is simply 50.
// Mode 2, relative mode, is similar to position mode: the parameter is interpreted as a position. Like position mode, parameters in relative mode can be read from or written to.  However, relative mode paramters count from a value called the relative base. The relative base starts at 0.

package main

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type IntComputer struct {
	inputs, outputs, state []int
	ip                     int // instruction pointer
	relativeBase           int
	terminated             bool
}

type ParamMode int

const (
	POS ParamMode = iota
	DIRECT
	RELATIVE
)

var computer_ram int = 20000

// Converts a computer state string into a slice of integer opcodes and parameters
func initComputer(state string, ins []int) (computer IntComputer) {
	log.Debugf("New computer: %#v", state)
	log.Debugf("Inputs: %#v", ins)

	computer.inputs = ins

	computer.ip = 0

	inst := strings.Split(state, ",")
	//computer.state = make([]int, len(inst))
	computer.state = make([]int, computer_ram)
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
	log.Debugf("opAdd (ip=%d)", computer.ip)
	pc := computer.ip
	p1 := getParamValue(pm[1], pc+1, computer)
	p2 := getParamValue(pm[2], pc+2, computer)
	var pdest int
	if pm[3] == RELATIVE {
		pdest = computer.relativeBase + computer.state[pc+3]
	} else {
		pdest = computer.state[pc+3]
	}
	log.Debugf("opAdd %d(@%d) + %d(@%d) into loc %d\n", p1, pc+1, p2, pc+2, pdest)
	computer.state[pdest] = p1 + p2
	computer.ip += 4
	return computer
}

func opMult(pm [4]ParamMode, computer IntComputer) IntComputer {
	log.Debugf("opMult (ip=%d)", computer.ip)
	pc := computer.ip
	p1 := getParamValue(pm[1], pc+1, computer)
	p2 := getParamValue(pm[2], pc+2, computer)
	var pdest int
	if pm[3] == RELATIVE {
		pdest = computer.relativeBase + computer.state[pc+3]
	} else {
		pdest = computer.state[pc+3]
	}
	log.Debugf("opMult %d(@%d) + %d(@%d) into %d\n", p1, pc+1, p2, pc+2, pc+3)
	computer.state[pdest] = p1 * p2
	computer.ip += 4
	return computer
}

func opReadInput(pm [4]ParamMode, computer IntComputer) IntComputer {
	log.Debugf("opReadInput (ip=%d)", computer.ip)
	pc := computer.ip
	var pdest int
	if pm[1] == RELATIVE {
		pdest = computer.relativeBase + computer.state[pc+1]
	} else {
		pdest = computer.state[pc+1]
	}

	// Try to read input
	var input int
	if len(computer.inputs) > 0 {
		// Pop an input value
		n_inputs := len(computer.inputs)
		input, computer.inputs = computer.inputs[n_inputs-1], computer.inputs[:n_inputs-1]
		log.Debugf("Read input value %d\n", input)
	}

	computer.state[pdest] = input
	computer.ip += 2
	return computer
}

func opWriteOutput(pm [4]ParamMode, computer IntComputer) IntComputer {
	log.Debugf("opWriteOutput (ip=%d)", computer.ip)
	pc := computer.ip
	p1 := getParamValue(pm[1], pc+1, computer)
	log.Debugf("Wrote output: %#v\n", p1)
	computer.outputs = append(computer.outputs, p1)
	computer.ip += 2
	return computer
}

func opJumpIfTrue(pm [4]ParamMode, computer IntComputer) IntComputer {
	log.Debugf("opJumpIfTrue (ip=%d)", computer.ip)
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
	log.Debugf("opJumpIfFalse (ip=%d)", computer.ip)
	pc := computer.ip
	p1 := getParamValue(pm[1], pc+1, computer)
	p2 := getParamValue(pm[2], pc+2, computer)
	log.Debugf("opJumpIfFalse p1=%d, p2=%d", p1, p2)
	if p1 == 0 {
		computer.ip = p2
	} else {
		computer.ip += 3
	}

	return computer
}

func opLessThan(pm [4]ParamMode, computer IntComputer) IntComputer {
	log.Debugf("opLessThan (ip=%d)", computer.ip)
	pc := computer.ip
	p1 := getParamValue(pm[1], pc+1, computer)
	p2 := getParamValue(pm[2], pc+2, computer)
	var pdest int
	if pm[3] == RELATIVE {
		pdest = computer.relativeBase + computer.state[pc+3]
	} else {
		pdest = computer.state[pc+3]
	}
	if p1 < p2 {
		computer.state[pdest] = 1
	} else {
		computer.state[pdest] = 0
	}
	computer.ip += 4

	return computer
}

func opEquals(pm [4]ParamMode, computer IntComputer) IntComputer {
	log.Debugf("opEquals (ip=%d)", computer.ip)
	pc := computer.ip
	p1 := getParamValue(pm[1], pc+1, computer)
	p2 := getParamValue(pm[2], pc+2, computer)
	var pdest int
	if pm[3] == RELATIVE {
		pdest = computer.relativeBase + computer.state[pc+3]
	} else {
		pdest = computer.state[pc+3]
	}
	if p1 == p2 {
		computer.state[pdest] = 1
	} else {
		computer.state[pdest] = 0
	}
	computer.ip += 4

	return computer
}

func opSetRelativeBase(pm [4]ParamMode, computer IntComputer) IntComputer {
	log.Debugf("opSetRelativeBase (ip=%d)", computer.ip)
	pc := computer.ip
	p1 := getParamValue(pm[1], pc+1, computer)
	log.Debugf("relativeBase = %d", p1)
	computer.relativeBase += p1
	computer.ip += 2

	return computer
}

func processOp(op int, computer IntComputer) IntComputer {
	log.Debugf("State: %#v", computer.state)

	opcode, paramModes := decodeOp(op)

	switch opcode {
	case 1:
		computer = opAdd(paramModes, computer)
	case 2:
		computer = opMult(paramModes, computer)
	case 3:
		computer = opReadInput(paramModes, computer)
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
	case 9:
		computer = opSetRelativeBase(paramModes, computer)
	default:
		log.Fatal("Unexpected opcode!")
	}

	return computer
}

func decodeOp(op int) (opcode int, paramModes [4]ParamMode) {
	opcode = op % 100

	paramModes[0] = 999 // Param 0 is not well-defined; only 1-3 are valid for first, second, third parameter
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
		log.Debug("POS param at ", loc)
		return computer.state[computer.state[loc]]
	case DIRECT:
		log.Debug("DIRECT param at ", loc)
		return computer.state[loc]
	case RELATIVE:
		log.Debug("RELATIVE param at ", computer.relativeBase+computer.state[loc])
		return computer.state[computer.relativeBase+computer.state[loc]]
	}
	log.Fatal("Unknown parameter mode ", mode)
	return 0
}
