package main

import (
	"strings"
	"testing"
)

const example = `16,1,2,0,4,2,7,1,2,14`

func TestExamplePart1(t *testing.T) {
	var lines [][]string
	for _, line := range strings.Split(example, "\n") {
		lines = append(lines, strings.Fields(line))
	}

	in := parseInput(lines)

	actual := solve(in, true)

	const expected = 37

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

	actual := solve(in, false)

	const expected = 168

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}
