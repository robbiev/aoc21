package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	input := []instruction{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	}

	position := processInstructions1(input)
	actual := position.horizontal * position.depth

	const expected = 150

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}

func TestExamplePart2(t *testing.T) {
	input := []instruction{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	}

	position := processInstructions2(input)
	actual := position.horizontal * position.depth

	const expected = 900

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}
