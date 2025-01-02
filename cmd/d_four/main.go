package main

import (
	"bufio"
	"fmt"
	"os"

	ceressearch "github.com/DAbraham2/aoc24/ceres_search"
)

func main() {
	matrix := ceressearch.ByteMatrix{}

	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		matrix = append(matrix, []byte(scanner.Text()))
	}

	xmas_count := ceressearch.Count(matrix)
	max_count := ceressearch.CountMas(matrix)

	fmt.Printf("Solution for D4 part 1: %d", xmas_count)
	fmt.Println()
	fmt.Printf("Solution for D4 part 2: %d", max_count)
	fmt.Println()
}
