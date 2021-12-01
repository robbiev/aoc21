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

	if len(os.Args) < 2 {
		log.Fatal("specify '1' or '2'")
	}
	if os.Args[1] == "1" {
		fmt.Println(numberOfTimesDepthIncreases(input))
	} else if os.Args[1] == "2" {
		fmt.Println(numberOfTimesDepthIncreasesSliding(input))
	} else {
		log.Fatal("specify '1' or '2'")
	}
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

func numberOfTimesDepthIncreasesSliding(input []int) int {
	if len(input) < 4 {
		return 0
	}

	lastThreeSum := input[0] + input[1] + input[2]

	var increases int
	for i := 3; i < len(input); i++ {
		a := lastThreeSum
		b := a - input[i-3] + input[i]
		lastThreeSum = b
		if b > a {
			increases++
		}
	}

	return increases
}
