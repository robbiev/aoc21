package main

import (
	"strings"
	"testing"
)

const example = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func TestExamplePart1(t *testing.T) {
	var lines [][]string
	for _, line := range strings.Split(example, "\n") {
		lines = append(lines, strings.Fields(line))
	}

	in := parseInput(lines)

	actual := part1(in, 10)

	const expected = 5

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}

func TestExamplePart2(t *testing.T) {
	var lines [][]string
	for _, line := range strings.Split(example, "\n") {
		lines = append(lines, strings.Fields(line))
	}

	in := parseInput2(lines)

	// in = []line{{
	// 	from: coord{9, 7},
	// 	to:   coord{7, 9},
	// 	from: coord{1, 1},
	// 	to:   coord{3, 3},
	// }}
	actual := part2(in, 10)

	const expected = 12

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}
