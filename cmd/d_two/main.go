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
	for scanner.Scan() {
		line_contents := strings.Split(scanner.Text(), " ")
		var input []int
		for _, v := range line_contents {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			input = append(input, i)
		}

		if rnreports.Safe(&input) {
			safe_lines++
		}
	}

	fmt.Printf("Day two solution for task1: %d", safe_lines)
	fmt.Println()
}
