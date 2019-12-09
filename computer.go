// Code related to the ship computer.
package main

import (
	"log"
	"strconv"
	"strings"
)

type ParamMode int

const (
	POS ParamMode = iota
	DIRECT
)

var input, output int

// Converts a computer state string into a slice of integer opcodes and parameters
func initComputer(state string, computerInput int) []int {
	input = computerInput

	inst := strings.Split(state, ",")
	comp := make([]int, len(inst))
	for id, item := range inst {
		opcode, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal("Can't convert instruction to int")
		}
		comp[id] = opcode
	}
	return comp
}

// Return current state of computer as a string
func snapshotComputer(comp []int) (int, string) {
	strCode := make([]string, len(comp))
	for id, item := range comp {
		strCode[id] = strconv.Itoa(item)
	}
	return output, strings.Join(strCode, ",")
}

// Given a computer state and instruction pointer, run the program
func run(computer []int, pc int) []int {
	//log.Printf("In: %d, Out: %d, Computer: %#v", input, output, computer)
	op := computer[pc]
	if op == 99 {
		return computer
	}

	nextIP := processOp(op, pc, computer)

	updatedState := run(computer, nextIP)

	return updatedState
}

func opAdd(pm [4]ParamMode, pc int, computer []int) (nextIP int) {
	p1 := getParamValue(pm[1], pc+1, computer)
	p2 := getParamValue(pm[2], pc+2, computer)
	pdest := computer[pc+3]
	nextIP = pc + 4
	computer[pdest] = p1 + p2

	return
}

func opMult(pm [4]ParamMode, pc int, computer []int) (nextIP int) {
	p1 := getParamValue(pm[1], pc+1, computer)
	p2 := getParamValue(pm[2], pc+2, computer)
	pdest := computer[pc+3]
	nextIP = pc + 4
	computer[pdest] = p1 * p2

	return
}

func opReadInput(pc int, computer []int) (nextIP int) {
	pdest := computer[pc+1]
	nextIP = pc + 2
	computer[pdest] = input

	return
}

func opWriteOutput(pm [4]ParamMode, pc int, computer []int) (nextIP int) {
	p1 := getParamValue(pm[1], pc+1, computer)
	nextIP = pc + 2
	output = p1

	return
}

func opJumpIfTrue(pm [4]ParamMode, pc int, computer []int) (nextIP int) {
	p1 := getParamValue(pm[1], pc+1, computer)
	p2 := getParamValue(pm[2], pc+2, computer)
	if p1 != 0 {
		nextIP = p2
	} else {
		nextIP = pc + 3
	}

	return
}

func opJumpIfFalse(pm [4]ParamMode, pc int, computer []int) (nextIP int) {
	p1 := getParamValue(pm[1], pc+1, computer)
	p2 := getParamValue(pm[2], pc+2, computer)
	if p1 == 0 {
		nextIP = p2
	} else {
		nextIP = pc + 3
	}

	return
}

func opLessThan(pm [4]ParamMode, pc int, computer []int) (nextIP int) {
	p1 := getParamValue(pm[1], pc+1, computer)
	p2 := getParamValue(pm[2], pc+2, computer)
	pdest := computer[pc+3]
	if p1 < p2 {
		computer[pdest] = 1
	} else {
		computer[pdest] = 0
	}
	nextIP = pc + 4

	return
}

func opEquals(pm [4]ParamMode, pc int, computer []int) (nextIP int) {
	p1 := getParamValue(pm[1], pc+1, computer)
	p2 := getParamValue(pm[2], pc+2, computer)
	pdest := computer[pc+3]
	if p1 == p2 {
		computer[pdest] = 1
	} else {
		computer[pdest] = 0
	}
	nextIP = pc + 4

	return
}

func processOp(op, pc int, computer []int) (nextIP int) {
	opcode, paramModes := decodeOp(op)

	switch opcode {
	case 1:
		nextIP = opAdd(paramModes, pc, computer)
	case 2:
		nextIP = opMult(paramModes, pc, computer)
	case 3:
		nextIP = opReadInput(pc, computer)
	case 4:
		nextIP = opWriteOutput(paramModes, pc, computer)
	case 5:
		nextIP = opJumpIfTrue(paramModes, pc, computer)
	case 6:
		nextIP = opJumpIfFalse(paramModes, pc, computer)
	case 7:
		nextIP = opLessThan(paramModes, pc, computer)
	case 8:
		nextIP = opEquals(paramModes, pc, computer)
	default:
		log.Fatal("Unexpected opcode!")
	}

	return nextIP
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

func getParamValue(mode ParamMode, loc int, computer []int) int {
	switch mode {
	case POS:
		return computer[computer[loc]]
	case DIRECT:
		return computer[loc]
	}
	log.Fatal("Unknown parameter mode ", mode)
	return 0
}
