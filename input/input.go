package input

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func MustSlurp(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var input [][]string
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		input = append(input, split)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func MustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
