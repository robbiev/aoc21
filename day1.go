package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var input []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(numberOfTimesDepthIncreases(input))
}

func numberOfTimesDepthIncreases(input []int) int {
	if len(input) <= 1 {
		return 0
	}

	var increases int
	for i := 1; i < len(input); i++ {
		previous := input[i-1]
		current := input[i]
		if current > previous {
			increases++
		}
	}

	return increases
}
