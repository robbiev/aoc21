package main

import (
	"strings"
	"testing"
)

const example = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestExamplePart1(t *testing.T) {
	var lines [][]string
	for _, line := range strings.Split(example, "\n") {
		lines = append(lines, strings.Fields(line))
	}

	in := parseInput(lines)

	actual := part1(in)

	const expected = 15

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}

func TestExamplePart2(t *testing.T) {
	var lines [][]string
	for _, line := range strings.Split(example, "\n") {
		lines = append(lines, strings.Fields(line))
	}

	in := parseInput(lines)

	actual := part2(in)

	const expected = 1134

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}
