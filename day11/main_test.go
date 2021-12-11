package main

import (
	"strings"
	"testing"
)

const example1 = `11111
19991
19191
19991
11111`

const example2 = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

func TestExamplePart1(t *testing.T) {
	var lines [][]string
	for _, line := range strings.Split(example2, "\n") {
		lines = append(lines, strings.Fields(line))
	}

	in := parseInput(lines)

	actual := part1(in, 100).flashes

	const expected = 1656

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}

func TestExamplePart2(t *testing.T) {
	var lines [][]string
	for _, line := range strings.Split(example2, "\n") {
		lines = append(lines, strings.Fields(line))
	}

	in := parseInput(lines)

	actual := part1(in, 200).firstAllFlashStep

	const expected = 195

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}
