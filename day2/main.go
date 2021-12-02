package main

import (
	"fmt"
	"log"
	"os"

	"github.com/robbiev/aoc21/input"
)

type position struct {
	horizontal int
	depth      int
	aim        int
}

type instruction struct {
	command string
	units   int
}

func processInstructions1(instructions []instruction) position {
	var pos position
	for _, instr := range instructions {
		switch instr.command {
		case "forward":
			pos.horizontal += instr.units
		case "down":
			pos.depth += instr.units
		case "up":
			pos.depth -= instr.units
		}
	}
	return pos
}

func processInstructions2(instructions []instruction) position {
	var pos position
	for _, instr := range instructions {
		switch instr.command {
		case "forward":
			pos.horizontal += instr.units
			pos.depth += (pos.aim * instr.units)
		case "down":
			pos.aim += instr.units
		case "up":
			pos.aim -= instr.units
		}
	}
	return pos
}

func calculateResult(pos position) int {
	return pos.horizontal * pos.depth
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("specify '1' or '2'")
	}

	var instr []instruction
	for _, line := range input.MustSlurp("input.txt") {
		instr = append(instr, instruction{
			command: line[0],
			units:   input.MustParseInt(line[1]),
		})
	}

	switch os.Args[1] {
	case "1":
		// 2027977
		fmt.Println(calculateResult(processInstructions1(instr)))
	case "2":
		// 1903644897
		fmt.Println(calculateResult(processInstructions2(instr)))
	default:
		log.Fatal("specify '1' or '2'")
	}
}
