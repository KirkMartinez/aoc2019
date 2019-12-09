package main

import (
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
	log.Println("Day 1 prelim tests passed.")
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
	log.Printf("Day 1, part 1 solution: %d\n", totalFuel)
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

	log.Printf("Day 1, part 2 solution: %d\n", total)
}

func main() {
	prelimTests()
	part1()
	part2()
}
