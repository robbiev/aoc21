package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/robbiev/aoc21/input"
)

type result struct {
	illegalScore int
	missingScore int
}

var openToClose = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var scoreIllegal = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var scoreIncomplete = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func solve(lines [][]string) result {
	var illegals []string
	var missing [][]string

NextLine:
	for _, line := range lines {
		var stack []string
		for _, char := range line {
			switch char {
			case "(", "[", "{", "<":
				stack = append(stack, char)
			default:
				opening := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if char != openToClose[opening] {
					// corrupted
					illegals = append(illegals, char)
					continue NextLine
				}
			}
		}
		if len(stack) > 0 {
			// incomplete

			// the order matters for scoring
			var fixups []string
			for i := len(stack) - 1; i >= 0; i-- {
				opening := stack[i]
				fixups = append(fixups, openToClose[opening])
			}
			missing = append(missing, fixups)
		}
	}

	var illegalScore int
	for _, illegal := range illegals {
		illegalScore += scoreIllegal[illegal]
	}

	var missingScore int
	{
		var missingScores []int
		for _, miss := range missing {
			var missingScore int
			for _, ch := range miss {
				missingScore *= 5
				missingScore += scoreIncomplete[ch]
			}
			missingScores = append(missingScores, missingScore)
		}
		sort.Ints(missingScores)
		missingScore = missingScores[len(missingScores)/2]
	}

	return result{
		illegalScore: illegalScore,
		missingScore: missingScore,
	}
}

func parseInput(lines [][]string) [][]string {
	var out [][]string
	for _, line := range lines {
		nums := strings.Split(line[0], "")
		var delims []string
		for _, n := range nums {
			delims = append(delims, n)
		}
		out = append(out, delims)
	}
	return out
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("specify '1' or '2'")
	}
	lines := input.MustSlurp("input.txt")
	in := parseInput(lines)

	switch os.Args[1] {
	case "1":
		// 323613
		fmt.Println(solve(in).illegalScore)
	case "2":
		// 3103006161
		fmt.Println(solve(in).missingScore)
	default:
		log.Fatal("specify '1' or '2'")
	}
}
