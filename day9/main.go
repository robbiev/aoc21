package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/robbiev/aoc21/input"
)

type coord struct {
	row, col int
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
		{row: r - 1, col: c}, // up
		{row: r + 1, col: c}, // down
		{row: r, col: c - 1}, // left
		{row: r, col: c + 1}, // right
	}
}

func filterValid(coords []coord, maxRow, maxCol int) []coord {
	var result []coord
	for _, coord := range coords {
		if coord.row != clamp(coord.row, 0, maxRow) || coord.col != clamp(coord.col, 0, maxCol) {
			continue
		}
		result = append(result, coord)
	}
	return result
}

func smolNums(grid [][]int) []coord {
	var smolNums []coord
	for r, row := range grid {
		for c, _ := range row {
			number := grid[r][c]
			neighbours := adjacent(r, c)
			neighbours = filterValid(neighbours, len(grid)-1, len(grid[0])-1)
			smol := true
			for _, n := range neighbours {
				neighbour := grid[n.row][n.col]
				smol = smol && number < neighbour
			}
			if smol {
				smolNums = append(smolNums, coord{row: r, col: c})
			}
		}
	}
	return smolNums
}

func part1(grid [][]int) int {
	smoln := smolNums(grid)
	var totalRiskLevel int
	for _, smoln := range smoln {
		riskLevel := grid[smoln.row][smoln.col] + 1
		totalRiskLevel += riskLevel
	}
	return totalRiskLevel
}

func part2(grid [][]int) int {
	basins := smolNums(grid)
	var basinSizes []int
	for _, basin := range basins {
		basinSize := 1
		explore := []coord{basin}
		discovered := map[coord]bool{basin: true}
		for len(explore) > 0 {
			coord := explore[0]
			explore = explore[1:]
			neighbours := adjacent(coord.row, coord.col)
			neighbours = filterValid(neighbours, len(grid)-1, len(grid[0])-1)
			var basinBuddies int
			for _, n := range neighbours {
				if discovered[n] {
					continue
				}
				discovered[n] = true
				neighbour := grid[n.row][n.col]
				if neighbour == 9 {
					continue
				}
				explore = append(explore, n)
				basinBuddies++
			}
			basinSize += basinBuddies
		}

		basinSizes = append(basinSizes, basinSize)
	}

	sort.Slice(basinSizes, func(i, j int) bool {
		return basinSizes[i] > basinSizes[j]
	})
	product := 1
	for _, size := range basinSizes[:3] {
		product *= size
	}
	return product
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
		// 530
		fmt.Println(part1(in))
	case "2":
		// 1019494
		fmt.Println(part2(in))
	default:
		log.Fatal("specify '1' or '2'")
	}
}
