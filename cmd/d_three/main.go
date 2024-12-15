package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	data := string(f)

	re := regexp.MustCompile(`mul\(([1-9][0-9]{0,2}),([1-9][0-9]{0,2})\)|do\(\)|don\'t\(\)`)

	res := re.FindAllStringSubmatch(data, -1)
	enabled := true
	prod := 0
	prod_t2 := 0
	for _, v := range res {
		switch v[0] {
		case "do()":
			enabled = true
			break
		case "don't()":
			enabled = false
			break
		default:
			r, _ := strconv.Atoi(v[1])
			l, _ := strconv.Atoi(v[2])

			prod += r * l
			if enabled {
				prod_t2 += r * l
			}
		}
	}

	fmt.Printf("Day three solution for part1: %d", prod)
	fmt.Println()
	fmt.Printf("Day three solution for part2: %d", prod_t2)
	fmt.Println()
}
