package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/robbiev/aoc21/input"
)

func part1(input []line, size int) int {
	overlapCounts := map[coord]int{}
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			for _, line := range input {
				var overlap bool

				var fromx, tox, fromy, toy int

				fromx, tox = line.from.x, line.to.x
				if line.from.x > line.to.x {
					fromx, tox = tox, fromx
				}
				fromy, toy = line.from.y, line.to.y
				if line.from.y > line.to.y {
					fromy, toy = toy, fromy
				}

				for x2 := fromx; x2 <= tox; x2++ {
					for y2 := fromy; y2 <= toy; y2++ {
						if x2 == x && y2 == y {
							overlap = true
						}
					}
				}

				if overlap {
					overlapCounts[coord{x, y}] = overlapCounts[coord{x, y}] + 1
				}
			}
		}
	}

	var atLeastTwoCount int
	for _, count := range overlapCounts {
		if count >= 2 {
			atLeastTwoCount++
		}
	}
	return atLeastTwoCount
}

func part2(input []line, size int) int {
	overlapCounts := map[coord]int{}
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			for _, line := range input {
				var overlap bool

				if line.from.x == line.to.x || line.from.y == line.to.y {
					var fromx, tox, fromy, toy int

					fromx, tox = line.from.x, line.to.x
					if line.from.x > line.to.x {
						fromx, tox = tox, fromx
					}
					fromy, toy = line.from.y, line.to.y
					if line.from.y > line.to.y {
						fromy, toy = toy, fromy
					}

					for x2 := fromx; x2 <= tox; x2++ {
						for y2 := fromy; y2 <= toy; y2++ {
							if x2 == x && y2 == y {
								overlap = true
							}
						}
					}
				} else {

					xdiff, ydiff := 1, 1
					if line.from.x > line.to.x {
						xdiff = -1
					}
					if line.from.y > line.to.y {
						ydiff = -1
					}
					count := line.to.x - line.from.x
					if count < 0 {
						count = -count
					}

					x2, y2 := line.from.x, line.from.y
					for ii := 0; ii <= count; ii++ {
						if x2 == x && y2 == y {
							overlap = true
						}
						if y2 == line.to.y {
							break
						}
						if x2 == line.to.x {
							break
						}
						x2 += xdiff
						y2 += ydiff
					}
				}

				if overlap {
					overlapCounts[coord{x, y}] = overlapCounts[coord{x, y}] + 1
				}
			}
		}
	}

	var atLeastTwoCount int
	for _, count := range overlapCounts {
		if count >= 2 {
			atLeastTwoCount++
		}
	}
	return atLeastTwoCount
}

type line struct {
	from coord
	to   coord
}

type coord struct {
	x, y int
}

func parseInput(lines [][]string) []line {
	var in []line
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
			in = append(in, line{from, to})
		}

	}
	return in
}

func parseInput2(lines [][]string) []line {
	var in []line
	for _, lin := range lines {
		var from coord
		fromStr := strings.Split(lin[0], ",")
		from.x = input.MustParseInt(fromStr[0])
		from.y = input.MustParseInt(fromStr[1])

		var to coord
		toStr := strings.Split(lin[2], ",")
		to.x = input.MustParseInt(toStr[0])
		to.y = input.MustParseInt(toStr[1])

		in = append(in, line{from, to})
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
		fmt.Println(part1(in, 1000))
	case "2":
		in = parseInput2(lines)
		fmt.Println(part2(in, 1000))
	default:
		log.Fatal("specify '1' or '2'")
	}
}
