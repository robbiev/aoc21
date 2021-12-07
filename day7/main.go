package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/robbiev/aoc21/input"
)

const (
	maxUint = ^uint(0)
	maxInt  = int(maxUint >> 1)
)

func solve(crabs []int, constantCost bool) int {
	if len(crabs) < 2 {
		return 0
	}
	min, max := crabs[0], crabs[0]
	for _, c := range crabs[1:] {
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}

	minFuelCost := maxInt
	for i := min; i <= max; i++ {
		var fuelCost int
		for _, c := range crabs {
			diff := c - i
			if diff < 0 {
				diff = -diff
			}
			cost := diff
			if !constantCost {
				cost = (diff + 1) * diff / 2
			}

			fuelCost += cost
		}

		if fuelCost < minFuelCost {
			minFuelCost = fuelCost
		}
	}

	return minFuelCost
}

func parseInput(lines [][]string) []int {
	var crabs []int
	for _, line := range lines {
		for _, number := range strings.Split(line[0], ",") {
			crabs = append(crabs, input.MustParseInt(number))
		}
	}
	return crabs
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("specify '1' or '2'")
	}
	lines := input.MustSlurp("input.txt")
	in := parseInput(lines)

	switch os.Args[1] {
	case "1":
		// 347509
		fmt.Println(solve(in, true))
	case "2":
		// 98257206
		fmt.Println(solve(in, false))
	default:
		log.Fatal("specify '1' or '2'")
	}
}
