package main

import (
	"log"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type WireInfo struct {
	wire_bitfield int
	step_count    [2]int // there are only two wires
}

func part1PrelimTests() {
	type Test struct {
		input  []string
		output int
	}

	tests := [3]Test{
		Test{[]string{"R8,U5,L5,D3",
			"U7,R6,D4,L4"}, 6},
		Test{[]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72",
			"U62,R66,U55,R34,D71,R55,D58,R83"}, 159},
		Test{[]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, 135}}
	for id, test := range tests {
		result := runPart1(test.input)
		if result != test.output {
			log.Println("Failed test #" + strconv.Itoa(id))
			log.Printf("Expected %d got %#v", test.output, result)
			log.Fatal("Stopping.")
		}
	}
	log.Println("Day 3, part 1 prelim tests passed.")
}

func part2PrelimTests() {
	type Test struct {
		input  []string
		output int
	}

	tests := [3]Test{
		Test{[]string{"R8,U5,L5,D3",
			"U7,R6,D4,L4"}, 30},
		Test{[]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72",
			"U62,R66,U55,R34,D71,R55,D58,R83"}, 610},
		Test{[]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, 410}}
	for id, test := range tests {
		result := runPart2(test.input)
		if result != test.output {
			log.Println("Failed test #" + strconv.Itoa(id))
			log.Printf("Expected %d got %#v", test.output, result)
			log.Fatal("Stopping.")
		}
	}
	log.Println("Day 3, part 2 prelim tests passed.")
}

func runPart1(wires []string) int {
	occupied := make(map[Point]int) // this int is a wire id bitfield

	var xPos, yPos int

	for id, wire := range wires {
		//log.Printf("Wire #%d: %s", id, wire)
		shape := strings.Split(wire, ",")
		for _, mvmt := range shape {
			xPos, yPos = move(occupied, xPos, yPos, mvmt, id)
		}
		xPos = 0
		yPos = 0
	}

	var crossings []Point

	for point, wireField := range occupied {
		// wire 0 will set bit 0, and wire 1, bit 1, so field will >2 iff both are set
		if wireField > 2 {
			crossings = append(crossings, point)
		}
	}

	var closest = dist(crossings[0])
	for _, crossing := range crossings {
		distance := dist(crossing)
		if distance < closest {
			closest = distance
		}
	}

	return closest
}

func runPart2(wires []string) int {
	occupied := make(map[Point]WireInfo)

	var xPos, yPos, wire_steps int

	for id, wire := range wires {
		//log.Printf("Wire #%d: %s", id, wire)
		shape := strings.Split(wire, ",")
		for _, mvmt := range shape {
			xPos, yPos, wire_steps = movePart2(occupied, xPos, yPos, mvmt, id, wire_steps)
		}
		xPos = 0
		yPos = 0
		wire_steps = 0
	}

	var crossings []Point

	for point, wireInfo := range occupied {
		// wire 0 will set bit 0, and wire 1, bit 1, so field will >2 iff both are set
		if wireInfo.wire_bitfield > 2 {
			crossings = append(crossings, point)
		}
	}

	//log.Println("Crossings: %#v", crossings)
	//for _, cross := range crossings {
	//  log.Println("WireInfo for %#v: %#v", cross, occupied[cross])
	//}

	first_crossing := crossings[0]
	var closest = occupied[first_crossing].step_count[0] + occupied[first_crossing].step_count[1]

	for _, crossing := range crossings {
		total_path_length := occupied[crossing].step_count[0] + occupied[crossing].step_count[1]
		if total_path_length < closest {
			closest = total_path_length
		}
	}

	return closest
}

func moveOffsets(mvmt string) (xOff, yOff int) {
	if mvmt[0] == 'U' {
		xOff, yOff = 0, -1
	}
	if mvmt[0] == 'D' {
		xOff, yOff = 0, 1
	}
	if mvmt[0] == 'L' {
		xOff, yOff = -1, 0
	}
	if mvmt[0] == 'R' {
		xOff, yOff = 1, 0
	}
	return xOff, yOff
}

func move(occupied map[Point]int, x int, y int, mvmt string, wire int) (int, int) {
	xOff, yOff := moveOffsets(mvmt)
	steps, err := strconv.Atoi(mvmt[1:])
	if err != nil {
		log.Fatal("Can't convert step count to int")
	}
	for i := 1; i <= steps; i++ {
		// Set the bit for this wire
		occupied[Point{x + xOff*i, y + yOff*i}] |= int(1 << uint(wire))
	}
	return x + xOff*steps, y + yOff*steps
}

func movePart2(occupied map[Point]WireInfo, x int, y int, mvmt string, wire, wire_steps int) (int, int, int) {
	xOff, yOff := moveOffsets(mvmt)
	steps, err := strconv.Atoi(mvmt[1:])
	if err != nil {
		log.Fatal("Can't convert step count to int")
	}
	for i := 1; i <= steps; i++ {
		// Set the bit for this wire
		this_point := Point{x + xOff*i, y + yOff*i}
		wire_info := WireInfo{occupied[this_point].wire_bitfield, occupied[this_point].step_count}
		wire_info.wire_bitfield |= int(1 << uint(wire))
		wire_steps++
		wire_info.step_count[wire] = wire_steps
		occupied[this_point] = wire_info
	}
	return x + xOff*steps, y + yOff*steps, wire_steps
}

func Abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func dist(pt Point) int {
	return Abs(pt.x) + Abs(pt.y)
}

func part1() {
	input := getData(3)
	result := runPart1(input)
	log.Printf("Day 3, part 1 solution: %d", result)
}

func part2() {
	input := getData(3)
	result := runPart2(input)
	log.Printf("Day 3, part 2 solution: %d", result)
}

func main() {
	part1PrelimTests()
	part2PrelimTests()
	part1()
	part2()
}
