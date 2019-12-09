package main

import (
	"log"
	"strconv"
	"strings"
)

type Test struct {
	input  string
	output int
}

type OrbitMap map[string][]string

const day int = 6

func prelimTests() {
	prelimPart1()
	prelimPart2()
}

func prelimPart1() {
	tests := [...]Test{
		Test{`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`, 42}}
	for id, test := range tests {
		input := strings.Split(test.input, "\n")
		result := runPart1(input)
		if result != test.output {
			log.Println("Failed part 1 test #" + strconv.Itoa(id))
			log.Printf("Expected %d, but got %d", test.output, result)
			log.Fatal("Aborting.")
		}
	}
	log.Println("Day 6, part 1 prelim tests part.")
}

func prelimPart2() {
	tests := [...]Test{
		Test{`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`, 4}}
	for id, test := range tests {
		input := strings.Split(test.input, "\n")
		result := runPart2(input)
		if result != test.output {
			log.Println("Failed part 2 test #" + strconv.Itoa(id))
			log.Printf("Expected %d, but got %d", test.output, result)
			log.Fatal("Aborting.")
		}
	}
	log.Println("Day 6, part 2 prelim tests passed.")
}

func buildOrbitMap(orbits []string) (orbit_map OrbitMap) {
	orbit_map = make(OrbitMap)

	for _, orbit := range orbits {
		orbit_info := strings.Split(orbit, ")")
		orbited, orbited_by := orbit_info[0], orbit_info[1]
		orbit_map[orbited] = append(orbit_map[orbited], orbited_by)
	}
	return
}

func runPart1(orbits []string) (result int) {
	orbit_map := buildOrbitMap(orbits)

	var orbit_count int
	level := 1
	for current := orbit_map["COM"]; len(current) > 0; current = orbiting(current, orbit_map) {
		orbit_count += level * len(current)
		level += 1
	}

	return orbit_count
}
func runPart2(orbits []string) (result int) {
	orbit_map := buildOrbitMap(orbits)

	// Strategy:
	// Enumerate orbited objects starting from YOU and SAN
	// Stop when 1) an enumerate obj is in the other path, or 2) is SAN/YOU
	// Case 1) path is concatenation of those two paths terminated at the common point
	// Case 2) path is the single enumerated path to SAN/YOU

	// Strategy 2:
	// 2xBFS for YOU and SAN
	// Path is concat of both including only one common obj (the last)
	y := bfs("COM", "YOU", orbit_map, nil)
	s := bfs("COM", "SAN", orbit_map, nil)

	for i := 0; y[i+1] == s[i+1]; {
		y = y[1:]
		s = s[1:]
	}

	return len(y) + len(s) - 4
}

func bfs(current string, obj string, orbit_map OrbitMap, path_so_far []string) (path []string) {
	if current == obj {
		return path_so_far
	} else {
		if len(orbit_map[current]) == 0 {
			return nil
		} else {
			for _, orbiting := range orbit_map[current] {
				search := bfs(orbiting, obj, orbit_map, append(path_so_far, orbiting))
				if search != nil {
					return search
				}
			}
		}
	}
	return nil
}

func orbiting(objects []string, orbit_map OrbitMap) (orbiting []string) {
	for _, obj := range objects {
		orbiting = append(orbiting, orbit_map[obj]...)
	}
	return
}

func part1() {
	input := getData(day)
	result := runPart1(input)
	log.Printf("Day 6, part 1 solution: %d\n", result)
}

func part2() {
	input := getData(day)
	result := runPart2(input)
	log.Printf("Day 6, part 2 solution: %d\n", result)
}

func main() {
	prelimTests()
	part1()
	part2()
}
