package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/robbiev/aoc21/input"
)

func part1(fish []int) int {
	// each day, a 0 becomes a 6 and adds a new 8 to the end of the list
	// each other number decreases by 1 if it was present at the start of the day
	for i := 0; i < 80; i++ {
		var newFish int
		for i, f := range fish {
			if f == 0 {
				fish[i] = 6
				newFish++
			} else {
				fish[i]--
			}
		}
		for i := 0; i < newFish; i++ {
			fish = append(fish, 8)
		}
	}
	return len(fish)
}

func part2(in []int) int {
	var fish [9]int
	for _, f := range in {
		fish[f]++
	}
	for i := 0; i < 256; i++ {
		zeroes := fish[0]
		// minus one to all fish
		fish[0] = fish[1]
		fish[1] = fish[2]
		fish[2] = fish[3]
		fish[3] = fish[4]
		fish[4] = fish[5]
		fish[5] = fish[6]
		fish[6] = fish[7]
		fish[7] = fish[8]
		// for every 0 we add an 8
		fish[8] = zeroes
		// 0 becomes a 6
		fish[6] += zeroes
	}

	var sum int
	for _, f := range fish {
		sum += f
	}
	return sum
}

func parseInput(lines [][]string) []int {
	var fish []int
	for _, line := range lines {
		for _, number := range strings.Split(line[0], ",") {
			fish = append(fish, input.MustParseInt(number))
		}
	}
	return fish
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("specify '1' or '2'")
	}
	lines := input.MustSlurp("input.txt")
	in := parseInput(lines)

	switch os.Args[1] {
	case "1":
		// 390923
		fmt.Println(part1(in))
	case "2":
		// 1749945484935
		fmt.Println(part2(in))
	default:
		log.Fatal("specify '1' or '2'")
	}
}
