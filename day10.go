package main

import (
	"log"
	"math/big"
	"strconv"
	"strings"
)

const day int = 10

func prelimTests() {
	type Test struct {
		input           string
		x, y, asteroids int
	}

	tests := [...]Test{
		Test{`.#..#
.....
#####
....#
...##`, 3, 4, 8},
		Test{`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`, 5, 8, 33},
		Test{`#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`, 1, 2, 35},
		Test{`.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`, 6, 3, 41},
		Test{`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`, 11, 13, 210}}
	for id, test := range tests {
		lines := strings.Split(test.input, "\n")
		x, y, asteroids := runPart1(lines)
		if x != test.x || y != test.y || asteroids != test.asteroids {
			log.Printf("Expected (%d,%d) with %d asteroids, got (%d,%d) with %d asteroids", test.x, test.y, test.asteroids, x, y, asteroids)
			log.Fatal("Failed test #" + strconv.Itoa(id))
		}
	}
	log.Println("Prelim tests passed.")
}

func showMap(starmap []string) {
	for iy := 0; iy < len(starmap); iy++ {
		log.Printf("%s\n", starmap[iy])
	}
}

func runPart1(starmap []string) (x, y, asteroids int) {
	// Try every asteroid as an observatory
	// Find the largest asteroid count
	for ix := 0; ix < len(starmap[0]); ix++ {
		for iy := 0; iy < len(starmap); iy++ {
			// Observatories can only be on asteroids
			if starmap[iy][ix] == '#' {
				log.Printf("Testing observatory at %d,%d\n", ix, iy)
				showMap(starmap)
				log.Printf("Testing observatory at %d,%d\n", ix, iy)
				visibleAsteroids := countVisibleAsteroids(starmap, ix, iy)
				if visibleAsteroids > asteroids {
					x = ix
					y = iy
					asteroids = visibleAsteroids
				}
			}
		}
	}
	return
}

func duplicateStarmap(starmap []string) []string {
	cpy := make([]string, len(starmap))
	for i := 0; i < len(starmap); i++ {
		newString := make([]byte, len(starmap[i]))
		copy(newString, starmap[i])
		cpy[i] = string(newString)
	}
	return cpy
}

func gcd(a, b int) int {
	var result = new(big.Int)
	var bigA = new(big.Int).SetInt64(int64(a))
	var bigB = new(big.Int).SetInt64(int64(b))
	result.GCD(nil, nil, bigA, bigB)
	return int(result.Int64())
}

func countVisibleAsteroids(starmap []string, x, y int) (asteroids int) {
	visible := duplicateStarmap(starmap)
	// Remove any stars occluded by others when viewed from (x,y)
	for ix := 0; ix < len(visible[0]); ix++ {
		for iy := 0; iy < len(visible); iy++ {
			if starmap[iy][ix] == '#' {
				// Skip asteroid at location of observatory
				if x == ix && y == iy {
					continue
				} else {
					log.Printf("Star at %d,%d eliminates the following:\n", ix, iy)
				}
				offX := ix - x
				offY := iy - y
				// loop until we're off the grid
				for m := 1; true; m++ {
					scale := gcd(abs(offX), abs(offY))
					slopeX, slopeY := offX, offY
					if scale != 0 {
						slopeX, slopeY = slopeX/scale, slopeY/scale
					} else {
						// Horizontal and vertical block all stars behind them regardless of distance
						if slopeX == 0 {
							slopeY = slopeY / abs(slopeY)
						} else if slopeY == 0 {
							slopeX = slopeX / abs(slopeX)
						} else if abs(slopeX) == abs(slopeY) {
							slopeX, slopeY = slopeX/abs(slopeX), slopeY/abs(slopeY)
						}
					}
					wipeX, wipeY := ix+m*slopeX, iy+m*slopeY
					log.Printf("-- Candidate to remove %d,%d\n", wipeX, wipeY)
					if wipeX >= 0 && wipeX < len(visible[0]) && wipeY >= 0 && wipeY < len(visible) {
						log.Printf("Removing star at %d,%d\n", wipeX, wipeY)
						row := visible[wipeY]
						visible[wipeY] = row[:wipeX] + "." + row[wipeX+1:]
					} else {
						break // Outside bounds of the starmap
					}
				}
			}
		}
	}
	log.Printf("Visible map:\n")
	showMap(visible)
	return countStars(visible) - 1 // Don't could the observatory star
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func countStars(starmap []string) (result int) {
	for ix := 0; ix < len(starmap[0]); ix++ {
		for iy := 0; iy < len(starmap); iy++ {
			if starmap[iy][ix] == '#' {
				result++
			}
		}
	}
	return result
}

func part1() {
	input := getData(day)
	x, y, asteroids := runPart1(input)
	log.Printf("Day 10, part 1 solution: (%d,%d) with %d asteriods", x, y, asteroids)
}

func part2() {
}

func main() {
	prelimTests()
	part1()
	part2()
}
