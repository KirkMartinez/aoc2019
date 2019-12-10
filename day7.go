package main

import (
	prmt "github.com/gitchander/permutation"
	"log"
	"strconv"
)

type Test struct {
	program  string
	settings string
	output   int
}

const day int = 7

func prelimTests() {
	part1PrelimTests()
	part2PrelimTests()
}

func part1PrelimTests() {
	tests := [...]Test{
		Test{"3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0", "43210", 43210},
		Test{"3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0", "01234", 54321},
		Test{"3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0", "10432", 65210}}
	for id, test := range tests {
		max_thrust, settings := runPart1(test.program)
		if max_thrust != test.output || settings != test.settings {
			log.Println("Failed test #" + strconv.Itoa(id))
			log.Printf("Expected thrust %d with setting %s, got %d with setting %s", test.output, test.settings, max_thrust, settings)
			log.Fatal("Aborting.")
		}
	}
	log.Printf("Day %d, part 1 prelim tests passed.\n", day)
}

func part2PrelimTests() {
	tests := [...]Test{
		Test{"3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5", "98765", 139629729},
		Test{"3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10", "97856", 18216}}
	for id, test := range tests {
		max_thrust, settings := runPart2(test.program)
		if max_thrust != test.output || settings != test.settings {
			log.Println("Failed test #" + strconv.Itoa(id))
			log.Printf("Expected thrust %d with setting %s, got %d with setting %s", test.output, test.settings, max_thrust, settings)
			log.Fatal("Aborting.")
		}
	}
	log.Printf("Day %d, part 2 prelim tests passed.\n", day)
}

func runPart1(program string) (int, string) {
	var max_thrust int
	var best_setting string

	a_setting := []int{0, 1, 2, 3, 4}
	p := prmt.New(prmt.IntSlice(a_setting))

	for p.Next() {
		thr := thrust(program, a_setting)
		if thr > max_thrust {
			max_thrust = thr
			best_setting = ""
			for _, s := range a_setting {
				best_setting += strconv.Itoa(s)
			}
		}
	}
	return max_thrust, best_setting
}

func runPart2(program string) (int, string) {
	var max_thrust int
	var best_setting string

	a_setting := []int{9, 8, 7, 6, 5}
	p := prmt.New(prmt.IntSlice(a_setting))

	for p.Next() {
		thr := thrustPart2(program, a_setting)
		if thr > max_thrust {
			max_thrust = thr
			best_setting = ""
			for _, s := range a_setting {
				best_setting += strconv.Itoa(s)
			}
		}
	}
	return max_thrust, best_setting
}

func prependInput(computer IntComputer, input int) IntComputer {
	computer.inputs = append([]int{input}, computer.inputs...)
	return computer
}

func thrustPart2(program string, settings []int) (result int) {
	var amp [5]IntComputer
	var output int

	// Init amps with their settings
	for amp_id := 0; amp_id < 5; amp_id++ {
		if amp_id == 0 {
			// Amp A (index 0) gets initial input of zero
			amp[amp_id] = initComputer(program, []int{0, settings[amp_id]})
		} else {
			amp[amp_id] = initComputer(program, []int{settings[amp_id]})
		}
	}

	for {
	AMP_LOOP:
		for amp_id := 0; amp_id < 5; amp_id++ {
			for {
				if amp[amp_id].terminated {
					return output
				}
				//log.Printf("Executing step (ip=%d) for amp %d\n", amp[amp_id].ip, amp_id)
				amp[amp_id] = runStep(amp[amp_id])
				// Did this amp generate output?
				if len(amp[amp_id].outputs) > 0 {
					// Pop it off
					outcnt := len(amp[amp_id].outputs)
					output, amp[amp_id].outputs = amp[amp_id].outputs[outcnt-1], amp[amp_id].outputs[:outcnt-1]
					//log.Printf("Saw output %d from amp %d\n", output, amp_id)

					next_amp_id := amp_id + 1
					if next_amp_id == 5 {
						next_amp_id = 0
					}
					// There will only ever be one input stacked up here
					amp[next_amp_id] = prependInput(amp[next_amp_id], output)
					//log.Printf("Set amp %d to have inputs: %#v\n", next_amp_id, amp[next_amp_id].inputs)

					continue AMP_LOOP // Process the next amp in the sequence
				}
			}
		}
	}

	return 8675309 // Maybe Jenny can help debug?
}

func thrust(program string, settings []int) (result int) {
	var amp IntComputer
	var output int

	for amp_id := 0; amp_id < 5; amp_id++ {
		amp = initComputer(program, []int{output, settings[amp_id]})
		amp = run(amp)
		output = amp.outputs[len(amp.outputs)-1]
	}

	return output
}

func part1() {
	program := getData(day)
	thrust, setting := runPart1(program[0])
	log.Printf("Day 7, part 1 solution: max_thrust %d with setting: %s\n", thrust, setting)
}

func part2() {
	program := getData(day)
	thrust, setting := runPart2(program[0])
	log.Printf("Day 7, part 2 solution: max_thrust %d with setting: %s\n", thrust, setting)
}

func main() {
	prelimTests()
	part1()
	part2()
}
