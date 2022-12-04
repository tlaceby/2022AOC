package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tlaceby/go-utils/fs"
)

const INPUT_FILE = "input.txt"

func main() {
	println("AdventOfCode 2022-4\n")
	input, _ := fs.ReadTextFile(INPUT_FILE)

	fullyPairs := 0

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			break
		}

		pairs := strings.Split(line, ",")
		p1 := strings.Split(pairs[0], "-")
		p2 := strings.Split(pairs[1], "-")

		p1_start, _ := strconv.Atoi(p1[0])
		p1_end, _ := strconv.Atoi(p1[1])
		p2_start, _ := strconv.Atoi(p2[0])
		p2_end, _ := strconv.Atoi(p2[1])

		if p1_start <= p2_start && p1_end >= p2_start {
			fullyPairs += 1
		} else if p2_start <= p1_start && p2_end >= p1_start {
			fullyPairs += 1
		}

	}

	fmt.Printf("pairs %d\n", fullyPairs)
}
