package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/robbiev/aoc21/input"
)

func isLower(s string) bool {
	return s[0] >= 'a' && s[0] <= 'z'
}

func solve(m map[string][]string, revisitOneSmol bool) int {
	//fmt.Println(m)
	paths := visit(m, revisitOneSmol, "start", []string{"start"})
	return len(paths)
}

// A:[start c b end]
// b:[start A d end]
// c:[A] d:[b]
// end:[A b]
// start:[A b]
func visit(m map[string][]string, revisitOneSmol bool, elem string, path []string) [][]string {
	if elem == "end" {
		//fmt.Println("PATH", path)
		return [][]string{path}
	}

	var paths [][]string
	for _, friend := range m[elem] {
		if friend == "start" {
			continue
		}
		if isLower(friend) {
			lowerFreq := lowerPathFreq(path)
			if revisitOneSmol {
				var alreadyHaveLowerOccuringTwice bool
				for _, freq := range lowerFreq {
					if freq == 2 {
						alreadyHaveLowerOccuringTwice = true
					}
				}
				if alreadyHaveLowerOccuringTwice && lowerFreq[friend] >= 1 {
					continue
				}
			} else if lowerFreq[friend] == 1 {
				continue
			}
		}
		paths = append(paths, visit(m, revisitOneSmol, friend, append(path, friend))...)
	}
	return paths
}

func lowerPathFreq(path []string) map[string]int {
	freq := map[string]int{}
	for _, pe := range path {
		if isLower(pe) {
			freq[pe]++
		}
	}
	return freq
}

func parseInput(lines [][]string) map[string][]string {
	m := map[string][]string{}
	for _, line := range lines {
		nums := strings.Split(line[0], "-")
		m[nums[0]] = append(m[nums[0]], nums[1])
		m[nums[1]] = append(m[nums[1]], nums[0])
	}
	return m
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("specify '1' or '2'")
	}
	lines := input.MustSlurp("input.txt")
	in := parseInput(lines)

	switch os.Args[1] {
	case "1":
		// 5212
		fmt.Println(solve(in, false))
	case "2":
		// 134862
		fmt.Println(solve(in, true))
	default:
		log.Fatal("specify '1' or '2'")
	}
}
