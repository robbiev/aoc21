package main

import (
	"strings"
	"testing"
)

const example = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

func TestExamplePart1(t *testing.T) {
	var lines [][]string
	for _, line := range strings.Split(example, "\n") {
		lines = append(lines, strings.Fields(line))
	}

	in := parseInput(lines)

	actual := solve(in).illegalScore

	const expected = 26397

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

	actual := solve(in).missingScore

	const expected = 288957

	if expected != actual {
		t.Fatal("got", actual, "wanted", expected)
	}
}
