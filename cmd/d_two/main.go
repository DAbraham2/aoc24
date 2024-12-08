package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DAbraham2/aoc24/rnreports"
)

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	safe_lines := 0
	safe_lines_dampened := 0
	for scanner.Scan() {
		line_contents := strings.Split(scanner.Text(), " ")
		var input []int
		var input_cpy []int
		for _, v := range line_contents {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			input = append(input, i)
			input_cpy = append(input_cpy, i)
		}

		if rnreports.Safe(&input) {
			safe_lines++
		}

		if rnreports.SafeDampener(&input_cpy) {
			safe_lines_dampened++
		}
	}

	fmt.Printf("Day two solution for task1: %d", safe_lines)
	fmt.Println()
	fmt.Printf("Day two solution for task2: %d", safe_lines_dampened)
	fmt.Println()
}
