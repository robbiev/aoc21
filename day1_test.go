package main

import (
	"testing"
)

func TestDay1ExamplePart1(t *testing.T) {
	input := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}

	const expected = 7

	actual := numberOfTimesDepthIncreases(input)

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}

func TestDay1ExamplePart2(t *testing.T) {
	input := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}

	const expected = 5

	actual := numberOfTimesDepthIncreasesSliding(input)

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}
