package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/robbiev/aoc21/input"
)

type diagnostic struct {
	gammaRate   int
	epsilonRate int
}

type diagnostic2 struct {
	oxygenGenRating int
	co2ScrubRating  int
}

func processInstructions1(instructions []int, max int) diagnostic {
	var diag diagnostic

	for i := 0; i < max; i++ {
		var ones int

		shift := 1 << i
		for _, instr := range instructions {
			if instr&shift == shift {
				ones++
			}
		}
		zeroes := len(instructions) - ones
		if ones > zeroes {
			diag.gammaRate |= (1 << i)
			diag.epsilonRate &= ^(1 << i)
		} else {
			diag.gammaRate &= ^(1 << i)
			diag.epsilonRate |= (1 << i)
		}
	}

	return diag
}

func processInstructions2(instructions []int, max int) diagnostic2 {
	var diag diagnostic2
	diag.oxygenGenRating = selectDiagnostic(max-1, instructions, func(ones, zeroes int) bool {
		return ones >= zeroes
	})
	diag.co2ScrubRating = selectDiagnostic(max-1, instructions, func(ones, zeroes int) bool {
		return ones < zeroes
	})
	return diag
}

func selectDiagnostic(bitPos int, remaining []int, keepOnes func(ones, zeroes int) bool) int {
	if len(remaining) == 1 {
		return remaining[0]
	}
	if bitPos < 0 {
		panic("no answer")
	}

	shift := 1 << bitPos

	var ones int
	for _, rem := range remaining {
		if rem&shift == shift {
			ones++
		}
	}

	zeroes := len(remaining) - ones

	var bitToKeep int
	if keepOnes(ones, zeroes) {
		bitToKeep = shift
	}

	var include []int
	for _, rem := range remaining {
		if rem&shift == bitToKeep {
			include = append(include, rem)
		}
	}

	bitPos--
	return selectDiagnostic(bitPos, include, keepOnes)
}

func calculatePowerConsumption(diag diagnostic) int {
	return diag.gammaRate * diag.epsilonRate
}

func calculateLifeSupportRating(diag diagnostic2) int {
	return diag.oxygenGenRating * diag.co2ScrubRating
}

func mustParseBinaryInt(s string) int {
	v, err := strconv.ParseInt(s, 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	return int(v)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("specify '1' or '2'")
	}

	var instr []int
	for _, line := range input.MustSlurp("input.txt") {
		instr = append(instr, mustParseBinaryInt(line[0]))
	}

	switch os.Args[1] {
	case "1":
		// 2648450
		fmt.Println(calculatePowerConsumption(processInstructions1(instr, 12)))
	case "2":
		// 2845944
		fmt.Println(calculateLifeSupportRating(processInstructions2(instr, 12)))
	default:
		log.Fatal("specify '1' or '2'")
	}
}
