package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/robbiev/aoc21/input"
)

type point struct {
	x, y int
}

type paper struct {
	dots  []point
	folds []point
}

func foldX(x int, dot point) point {
	if dot.x < x {
		return dot // we fold left only
	}
	return point{
		x: dot.x - ((dot.x - x) * 2),
		y: dot.y,
	}
}

func foldY(y int, dot point) point {
	if dot.y < y {
		return dot // we fold up only
	}
	return point{
		x: dot.x,
		y: dot.y - ((dot.y - y) * 2),
	}
}

// y = up
// x = left
func solve(inputPaper paper, maxFolds int) int {
	dotMap := map[point]bool{}
	var bounds struct{ x, y int }
	for _, dot := range inputPaper.dots {
		dotMap[dot] = true
		if dot.x > bounds.x {
			bounds.x = dot.x
		}
		if dot.y > bounds.y {
			bounds.y = dot.y
		}
	}

	for _, fold := range inputPaper.folds[:maxFolds] {
		var foldFunc func(point) point
		if fold.x != 0 {
			foldFunc = func(dot point) point {
				return foldX(fold.x, dot)
			}
			bounds.x = bounds.x - fold.x - 1
		} else if fold.y != 0 {
			foldFunc = func(dot point) point {
				return foldY(fold.y, dot)
			}
			bounds.y = bounds.y - fold.y - 1
		}

		var replacements []point
		for dot, _ := range dotMap {
			newDot := foldFunc(dot)
			replacements = append(replacements, newDot)
		}
		for _, replace := range replacements {
			dotMap[replace] = true
		}
	}

	var dotCount int
	for dot, _ := range dotMap {
		if dot.x >= 0 && dot.x <= bounds.x && dot.y >= 0 && dot.y <= bounds.y {
			dotCount++
		}
	}

	for y := 0; y <= bounds.y; y++ {
		for x := 0; x <= bounds.x; x++ {
			if dotMap[point{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return dotCount
}

func parseInput(lines [][]string) paper {
	var in paper
	for i, dot := range lines {
		if len(dot) == 0 {
			// empty line, parse folds
			lines = lines[i+1:]
			break
		}
		xy := strings.Split(dot[0], ",")
		in.dots = append(in.dots, point{x: input.MustParseInt(xy[0]), y: input.MustParseInt(xy[1])})
	}
	for _, fold := range lines {
		split := strings.Split(fold[2], "=")
		var f point
		switch split[0] {
		case "x":
			f.x = input.MustParseInt(split[1])
		case "y":
			f.y = input.MustParseInt(split[1])
		}
		in.folds = append(in.folds, f)
	}
	return in
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("specify '1' or '2'")
	}
	lines := input.MustSlurp("input.txt")
	in := parseInput(lines)

	switch os.Args[1] {
	case "1":
		// 850
		fmt.Println(solve(in, 1))
	case "2":
		// AHGCPGAU
		fmt.Println(solve(in, len(in.folds)))
	default:
		log.Fatal("specify '1' or '2'")
	}
}
