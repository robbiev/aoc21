package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/robbiev/aoc21/input"
)

var standardAlphabet = []string{"a", "b", "c", "d", "e", "f", "g"}

type entry struct {
	signalPatterns []string // 10
	outputValues   []string // 4 digits
}

func part1(entries []entry) int {
	var count1478 int
	for _, entry := range entries {
		for _, outputVal := range entry.outputValues {
			switch len(outputVal) {
			// 1 = 2 seg
			// 7 = 3 seg
			// 4 = 4 seg
			// 8 = 7 seg
			case 2, 3, 4, 7:
				count1478++
			}
		}
	}
	return count1478
}

func without(chars []string, exclude string) []string {
	var n []string
	for _, ch := range chars {
		if ch == exclude {
			continue
		}
		n = append(n, ch)
	}
	return n
}

func permutations(str string, chars []string) []string {
	if len(chars) == 0 {
		return []string{str}
	}
	var perm []string
	for _, ch := range chars {
		newStr := str + ch
		newPerm := permutations(newStr, without(chars, ch))
		perm = append(perm, newPerm...)
	}
	return perm
}

//   aaaa
//  b    c
//  b    c
//   dddd
//  e    f
//  e    f
//   gggg

//  dddd
// e    a      deafgbc
// e    a
//  ffff
// g    b
// g    b
//  cccc
func part2(entries []entry) int {
	standardNumberRepr := map[string]int{
		"abcefg":  0,
		"cf":      1,
		"acdeg":   2,
		"acdfg":   3,
		"bcdf":    4,
		"abdfg":   5,
		"abdefg":  6,
		"acf":     7,
		"abcdefg": 8,
		"abcdfg":  9,
	}

	var sum int
	perms := permutations("", standardAlphabet)
	// perms := []string{
	// 	"deafgbc",
	// }
	for entryNum, entry := range entries {
		var foundSolution bool
		for _, alphabet := range perms {
			var validNumberCount int
			for _, pattern := range entry.signalPatterns {
				var maybeNumber []byte
				for i := 0; i < len(pattern); i++ {
					switch ch := pattern[i]; {
					case ch == alphabet[0]:
						maybeNumber = append(maybeNumber, 'a')
					case ch == alphabet[1]:
						maybeNumber = append(maybeNumber, 'b')
					case ch == alphabet[2]:
						maybeNumber = append(maybeNumber, 'c')
					case ch == alphabet[3]:
						maybeNumber = append(maybeNumber, 'd')
					case ch == alphabet[4]:
						maybeNumber = append(maybeNumber, 'e')
					case ch == alphabet[5]:
						maybeNumber = append(maybeNumber, 'f')
					case ch == alphabet[6]:
						maybeNumber = append(maybeNumber, 'g')
					}
				}
				sort.Slice(maybeNumber, func(i, j int) bool {
					return maybeNumber[i] < maybeNumber[j]
				})
				if _, ok := standardNumberRepr[string(maybeNumber)]; ok {
					validNumberCount++
				}
			}

			if validNumberCount == 10 {
				fmt.Println("found alphabet", entryNum)
				foundSolution = true
				var outNumbers string
				for _, pattern := range entry.outputValues {
					var maybeNumber []byte
					for i := 0; i < len(pattern); i++ {
						switch ch := pattern[i]; {
						case ch == alphabet[0]:
							maybeNumber = append(maybeNumber, 'a')
						case ch == alphabet[1]:
							maybeNumber = append(maybeNumber, 'b')
						case ch == alphabet[2]:
							maybeNumber = append(maybeNumber, 'c')
						case ch == alphabet[3]:
							maybeNumber = append(maybeNumber, 'd')
						case ch == alphabet[4]:
							maybeNumber = append(maybeNumber, 'e')
						case ch == alphabet[5]:
							maybeNumber = append(maybeNumber, 'f')
						case ch == alphabet[6]:
							maybeNumber = append(maybeNumber, 'g')
						}
					}
					sort.Slice(maybeNumber, func(i, j int) bool {
						return maybeNumber[i] < maybeNumber[j]
					})
					if num, ok := standardNumberRepr[string(maybeNumber)]; ok {
						outNumbers += strconv.Itoa(num)
					}
				}
				sum += input.MustParseInt(outNumbers)
			}
		}
		if !foundSolution {
			panic("no solution")
		}
	}
	return sum
}

func parseInput(lines [][]string) []entry {
	var entries []entry
	for _, line := range lines {
		var signalPatterns, outputValues []string
		for _, field := range line[:10] {
			signalPatterns = append(signalPatterns, field)
		}
		for _, field := range line[11:] {
			outputValues = append(outputValues, field)
		}
		entries = append(entries, entry{
			signalPatterns: signalPatterns,
			outputValues:   outputValues,
		})
	}
	return entries
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("specify '1' or '2'")
	}
	lines := input.MustSlurp("input.txt")
	in := parseInput(lines)

	switch os.Args[1] {
	case "1":
		// 278
		fmt.Println(part1(in))
	case "2":
		// 986179
		fmt.Println(part2(in))
	default:
		log.Fatal("specify '1' or '2'")
	}
}
