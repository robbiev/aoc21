package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/robbiev/aoc21/input"
)

type line struct {
	from coord
	to   coord
}

type coord struct {
	x, y int
}

func findOverlaps(lines [][]string, includeDiagonal bool) int {
	freq := map[coord]int{}
	for _, lin := range lines {
		var from coord
		fromStr := strings.Split(lin[0], ",")
		from.x = input.MustParseInt(fromStr[0])
		from.y = input.MustParseInt(fromStr[1])

		var to coord
		toStr := strings.Split(lin[2], ",")
		to.x = input.MustParseInt(toStr[0])
		to.y = input.MustParseInt(toStr[1])

		if from.x == to.x || from.y == to.y {
			fromx, tox := from.x, to.x
			if from.x > to.x {
				fromx, tox = tox, fromx
			}
			fromy, toy := from.y, to.y
			if from.y > to.y {
				fromy, toy = toy, fromy
			}

			for x := fromx; x <= tox; x++ {
				for y := fromy; y <= toy; y++ {
					freq[coord{x, y}] += 1
				}
			}
		} else if includeDiagonal {
			xstep, ystep := 1, 1
			if from.x > to.x {
				xstep = -1
			}
			if from.y > to.y {
				ystep = -1
			}
			count := to.x - from.x
			if count < 0 {
				count = -count
			}

			x, y := from.x, from.y
			for i := 0; i <= count; i++ {
				freq[coord{x, y}] += 1
				x += xstep
				y += ystep
			}
		}
	}

	var over2 int
	for _, count := range freq {
		if count >= 2 {
			over2++
		}
	}

	return over2
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("specify '1' or '2'")
	}
	lines := input.MustSlurp("input.txt")

	switch os.Args[1] {
	case "1":
		// 8622
		fmt.Println(findOverlaps(lines, false))
	case "2":
		// 22037
		fmt.Println(findOverlaps(lines, true))
	default:
		log.Fatal("specify '1' or '2'")
	}
}
