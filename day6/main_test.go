package main

import (
	"strings"
	"testing"
)

const example = `3,4,3,1,2`

func TestExamplePart1(t *testing.T) {
	var lines [][]string
	for _, line := range strings.Split(example, "\n") {
		lines = append(lines, strings.Fields(line))
	}

	in := parseInput(lines)

	actual := part1(in)

	const expected = 5934

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

	const expected = 26984457539

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}
