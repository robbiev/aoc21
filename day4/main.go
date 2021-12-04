package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/robbiev/aoc21/input"
)

const boardSize = 5

type board struct {
	numbers        [][]int
	rowNumbersSeen map[int]map[int]bool // row index to seen numbers
	colNumbersSeen map[int]map[int]bool // col index to seen numbers
}

type bingo struct {
	draw   []int
	boards []board
}

func part1(game bingo) int {
	for _, n := range game.draw {
		for _, b := range game.boards {
			for r := 0; r < boardSize; r++ {
				for c := 0; c < boardSize; c++ {
					if b.numbers[r][c] == n {
						if _, ok := b.rowNumbersSeen[r]; !ok {
							b.rowNumbersSeen[r] = map[int]bool{}
						}
						b.rowNumbersSeen[r][n] = true
						if _, ok := b.colNumbersSeen[c]; !ok {
							b.colNumbersSeen[c] = map[int]bool{}
						}
						b.colNumbersSeen[c][n] = true

						if len(b.colNumbersSeen[c]) == boardSize || len(b.rowNumbersSeen[r]) == boardSize {
							var sum int
							for r2 := 0; r2 < boardSize; r2++ {
								for c2 := 0; c2 < boardSize; c2++ {
									n := b.numbers[r2][c2]
									// it will appear in both rows and columns, so checking one is enough
									if _, ok := b.rowNumbersSeen[r2][n]; !ok {
										sum += n
									}
								}
							}
							return n * sum
						}
					}
				}
			}
		}
	}
	return 0 // no winner
}

func part2(game bingo) int {
	boardsWon := map[int]bool{}
	for _, n := range game.draw {
		for i, b := range game.boards {
			for r := 0; r < boardSize; r++ {
				for c := 0; c < boardSize; c++ {
					if b.numbers[r][c] == n {
						//fmt.Println("board", i+1, "matches", n)
						if _, ok := b.rowNumbersSeen[r]; !ok {
							b.rowNumbersSeen[r] = map[int]bool{}
						}
						b.rowNumbersSeen[r][n] = true
						if _, ok := b.colNumbersSeen[c]; !ok {
							b.colNumbersSeen[c] = map[int]bool{}
						}
						b.colNumbersSeen[c][n] = true

						if len(b.colNumbersSeen[c]) == boardSize || len(b.rowNumbersSeen[r]) == boardSize && !boardsWon[i] {
							boardsWon[i] = true
							// if it's the last board to win
							if len(boardsWon) == len(game.boards) {
								var sum int
								for r2 := 0; r2 < boardSize; r2++ {
									for c2 := 0; c2 < boardSize; c2++ {
										n := b.numbers[r2][c2]
										// it will appear in both rows and columns, so checking one is enough
										if _, ok := b.rowNumbersSeen[r2][n]; !ok {
											sum += n
										}
									}
								}
								return n * sum
							}
						}
					}
				}
			}
		}
	}
	return 0 // no winner
}

func parseInput(lines [][]string) bingo {
	var draw []int
	for _, number := range strings.Split(lines[0][0], ",") {
		draw = append(draw, input.MustParseInt(number))
	}
	//fmt.Println(draw)
	var boards []board
	var b *board
	for _, line := range lines[1:] {
		if len(line) == 0 {
			if b != nil {
				boards = append(boards, *b)
				//fmt.Println(b.numbers)
			}
			b = &board{
				numbers:        [][]int{},
				rowNumbersSeen: map[int]map[int]bool{},
				colNumbersSeen: map[int]map[int]bool{},
			}
			continue
		}

		var intLine []int
		for _, number := range line {
			intLine = append(intLine, input.MustParseInt(number))
		}
		b.numbers = append(b.numbers, intLine)
	}
	return bingo{
		draw:   draw,
		boards: boards,
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("specify '1' or '2'")
	}
	lines := input.MustSlurp("input.txt")
	game := parseInput(lines)

	switch os.Args[1] {
	case "1":
		// 8580
		fmt.Println(part1(game))
	case "2":
		// 9576
		fmt.Println(part2(game))
	default:
		log.Fatal("specify '1' or '2'")
	}
}
