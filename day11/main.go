package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/robbiev/aoc21/input"
)

const gridSize = 10

type coord struct {
	row, col int
}

type result struct {
	flashes           int
	firstAllFlashStep int
}

func clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func adjacent(r, c int) []coord {
	return []coord{
		{row: r - 1, col: c},     // up
		{row: r + 1, col: c},     // down
		{row: r, col: c - 1},     // left
		{row: r, col: c + 1},     // right
		{row: r - 1, col: c - 1}, // up left
		{row: r - 1, col: c + 1}, // up right
		{row: r + 1, col: c - 1}, // down left
		{row: r + 1, col: c + 1}, // down right
	}
}

func filterValid(coords []coord) []coord {
	var result []coord
	for _, coord := range coords {
		if coord.row != clamp(coord.row, 0, gridSize-1) || coord.col != clamp(coord.col, 0, gridSize-1) {
			continue
		}
		result = append(result, coord)
	}
	return result
}

func part1(grid [][]int, steps int) result {
	var out result
	for i := 0; i < steps; i++ {
		var flashGrid [10][10]bool

		// First, the energy level of each octopus increases by 1.
		for r, row := range grid {
			for c, _ := range row {
				grid[r][c]++
			}
		}

		// Then, any octopus with an energy level greater than 9 flashes.
		for r, row := range grid {
			for c, _ := range row {
				number := grid[r][c]

				if number > 9 && !flashGrid[r][c] {
					flashGrid[r][c] = true
					flashers := []coord{coord{row: r, col: c}}
					for len(flashers) > 0 {
						flasher := flashers[0]
						flashers = flashers[1:]
						neighbours := filterValid(adjacent(flasher.row, flasher.col))
						for _, n := range neighbours {
							// Flashing increases the energy level of all adjacent octopuses by
							// 1, including octopuses that are diagonally adjacent.
							grid[n.row][n.col]++

							// If this causes an octopus to have an energy level greater than 9,
							// it also flashes. This process continues as long as new octopuses
							// keep having their energy level increased beyond 9.
							if grid[n.row][n.col] <= 9 {
								continue
							}
							if flashGrid[n.row][n.col] {
								continue
							}
							// (An octopus can only flash at most once per step.)
							flashGrid[n.row][n.col] = true

							// Process flash
							flashers = append(flashers, n)
						}
					}
				}

			}
		}

		// Finally, any octopus that flashed during this step has its energy level
		// set to 0, as it used all of its energy to flash.
		flashesBefore := out.flashes
		for r, row := range flashGrid {
			for c, _ := range row {
				if flashGrid[r][c] {
					grid[r][c] = 0
					out.flashes++
				}
			}
		}

		// part 2
		if out.flashes-flashesBefore == (gridSize*gridSize) && out.firstAllFlashStep == 0 {
			out.firstAllFlashStep = i + 1
		}
	}

	return out
}

func part2(grid [][]int) int {
	return 0
}

func print(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func parseInput(lines [][]string) [][]int {
	var grid [][]int
	for _, line := range lines {
		nums := strings.Split(line[0], "")
		var numList []int
		for _, n := range nums {
			numList = append(numList, input.MustParseInt(n))
		}
		grid = append(grid, numList)
	}
	return grid
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("specify '1' or '2'")
	}
	lines := input.MustSlurp("input.txt")
	in := parseInput(lines)

	switch os.Args[1] {
	case "1":
		// 1571
		fmt.Println(part1(in, 100).flashes)
	case "2":
		// 387
		fmt.Println(part1(in, 1000).firstAllFlashStep)
	default:
		log.Fatal("specify '1' or '2'")
	}
}
