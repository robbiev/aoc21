package main

import (
	"strings"
	"testing"
)

// 10
const example1 = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

// 19
const example2 = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

// 226
const example3 = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`

func TestExamplePart1(t *testing.T) {
	var lines [][]string
	for _, line := range strings.Split(example1, "\n") {
		lines = append(lines, strings.Fields(line))
	}

	in := parseInput(lines)

	actual := solve(in, false)

	const expected = 10

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}

func TestExamplePart1b(t *testing.T) {
	var lines [][]string
	for _, line := range strings.Split(example2, "\n") {
		lines = append(lines, strings.Fields(line))
	}

	in := parseInput(lines)

	actual := solve(in, false)

	const expected = 19

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}

func TestExamplePart1c(t *testing.T) {
	var lines [][]string
	for _, line := range strings.Split(example3, "\n") {
		lines = append(lines, strings.Fields(line))
	}

	in := parseInput(lines)

	actual := solve(in, false)

	const expected = 226

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}

func TestExamplePart2(t *testing.T) {
	var lines [][]string
	for _, line := range strings.Split(example1, "\n") {
		lines = append(lines, strings.Fields(line))
	}

	in := parseInput(lines)

	actual := solve(in, true)

	const expected = 36

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}

func TestExamplePart2b(t *testing.T) {
	var lines [][]string
	for _, line := range strings.Split(example2, "\n") {
		lines = append(lines, strings.Fields(line))
	}

	in := parseInput(lines)

	actual := solve(in, true)

	const expected = 103

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}

func TestExamplePart2c(t *testing.T) {
	var lines [][]string
	for _, line := range strings.Split(example3, "\n") {
		lines = append(lines, strings.Fields(line))
	}

	in := parseInput(lines)

	actual := solve(in, true)

	const expected = 3509

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}
