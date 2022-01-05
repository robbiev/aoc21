package main

import (
	"strings"
	"testing"
)

const example1 = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

func TestFoldY(t *testing.T) {
	actual := foldY(7, point{x: 0, y: 14})

	expected := point{x: 0, y: 0}

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}

func TestExamplePart1(t *testing.T) {
	var lines [][]string
	for _, line := range strings.Split(example1, "\n") {
		lines = append(lines, strings.Fields(line))
	}

	in := parseInput(lines)

	actual := solve(in, 1)

	const expected = 17

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}
