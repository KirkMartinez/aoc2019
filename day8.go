package main

import (
	"log"
	"math"
)

type Layer [6][25]int

const day int = 8

func extractLayers(input string) (result []Layer) {
	var layer Layer

	for ch_index := 0; ch_index < len(input); {
		for row_id := 0; row_id < 6; row_id++ {
			for col_id := 0; col_id < 25; col_id++ {
				layer[row_id][col_id] = int(input[ch_index] - '0')
				ch_index++
			}
		}
		result = append(result, layer)
	}
	return result
}

func countZeros(layer Layer) (zeros int) {
	for row_id := 0; row_id < 6; row_id++ {
		for col_id := 0; col_id < 25; col_id++ {
			if layer[row_id][col_id] == 0 {
				zeros++
			}
		}
	}
	return zeros
}

func calcChecksum(layer Layer) (result int) {
	var ones, twos int
	for row_id := 0; row_id < 6; row_id++ {
		for col_id := 0; col_id < 25; col_id++ {
			switch layer[row_id][col_id] {
			case 1:
				ones++
			case 2:
				twos++
			}
		}
	}
	return ones * twos
}

func runPart1(input string) (result int) {
	layers := extractLayers(input)

	var fewestZerosLayer Layer
	var leastZerosSoFar int = math.MaxInt32
	for _, layer := range layers {
		zerosForThisLayer := countZeros(layer)
		if zerosForThisLayer < leastZerosSoFar {
			fewestZerosLayer = layer
			leastZerosSoFar = zerosForThisLayer
		}
	}

	return calcChecksum(fewestZerosLayer)
}

func part1() {
	input := getData(day)
	result := runPart1(input[0])
	log.Println("Day 8, part 1 solution: ", result)
}

func runPart2(input string) (result Layer) {
	layers := extractLayers(input)
	layer_cnt := len(layers)

	for row_id := 0; row_id < 6; row_id++ {
		for col_id := 0; col_id < 25; col_id++ {
			for layer_id := 0; layer_id < layer_cnt; layer_id++ {
				if layer_id == 0 {
					result[row_id][col_id] = layers[layer_id][row_id][col_id]
				} else {
					if result[row_id][col_id] == 2 {
						result[row_id][col_id] = layers[layer_id][row_id][col_id]
					}
				}
			}
		}
	}
	return result
}

func showLayer(layer Layer) {
	var line string

	for row_id := 0; row_id < 6; row_id++ {
		for col_id := 0; col_id < 25; col_id++ {
			if layer[row_id][col_id] == 0 {
				line += "X"
			} else {
				line += " "
			}
		}
		log.Println(line)
		line = ""
	}
}

func part2() {
	input := getData(day)
	result := runPart2(input[0])
	log.Print("Day 8, part 2 solution:")
	showLayer(result)
}

func main() {
	part1()
	part2()
}
