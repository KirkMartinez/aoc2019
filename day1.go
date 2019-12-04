package main

import (
	"fmt"
	"log"
	"strconv"
)

func fuel(mass int64) int64 {
	return int64(mass/3 - 2)
}

func prelimTests() {
	if fuel(12) != 2 {
		log.Fatal("wrong fuel val")
	}
	if fuel(14) != 2 {
		log.Fatal("wrong fuel val")
	}
	if fuel(1969) != 654 {
		log.Fatal("wrong fuel val")
	}
	if fuel(100756) != 33583 {
		log.Fatal("wrong fuel val")
	}
	fmt.Println("Prelim tests passed.")
}

func part1Fuel(masses []string) int64 {
	var total int64

	for _, line := range masses {
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("Failed to convert input mass to int")
		}
		total += fuel(int64(mass))
	}

	return total
}

func part1() {
	masses := getData(1)

	totalFuel := part1Fuel(masses)
	fmt.Println("Part 1 solution: " + strconv.FormatInt(totalFuel, 10))
}

func part2Fuel(masses []string) int64 {
	var total int64

	for _, line := range masses {
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("Failed to convert input mass to int")
		}
		total += recursiveFuel(int64(mass))
	}

	return total
}

func recursiveFuel(mass int64) int64 {
	var acc int64

	for v := fuel(mass); v > 0; v = fuel(v) {
		acc += v
	}
	return acc
}

func part2() {
	masses := getData(1)

	total := part2Fuel(masses)

	fmt.Println("Part 2 solution: " + strconv.FormatInt(total, 10))
}

func main() {
	prelimTests()
	part1()
	part2()
}
