package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	input := []int{
		mustParseBinaryInt("00100"),
		mustParseBinaryInt("11110"),
		mustParseBinaryInt("10110"),
		mustParseBinaryInt("10111"),
		mustParseBinaryInt("10101"),
		mustParseBinaryInt("01111"),
		mustParseBinaryInt("00111"),
		mustParseBinaryInt("11100"),
		mustParseBinaryInt("10000"),
		mustParseBinaryInt("11001"),
		mustParseBinaryInt("00010"),
		mustParseBinaryInt("01010"),
	}

	diag := processInstructions1(input, 5)
	actual := diag.gammaRate * diag.epsilonRate

	const expected = 198

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}

func TestExamplePart2(t *testing.T) {
	input := []int{
		mustParseBinaryInt("00100"),
		mustParseBinaryInt("11110"),
		mustParseBinaryInt("10110"),
		mustParseBinaryInt("10111"),
		mustParseBinaryInt("10101"),
		mustParseBinaryInt("01111"),
		mustParseBinaryInt("00111"),
		mustParseBinaryInt("11100"),
		mustParseBinaryInt("10000"),
		mustParseBinaryInt("11001"),
		mustParseBinaryInt("00010"),
		mustParseBinaryInt("01010"),
	}

	diag := processInstructions2(input, 5)
	actual := diag.oxygenGenRating * diag.co2ScrubRating

	const expected = 230

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}
