package main

import (
	"fmt"
	"strings"

	"github.com/tlaceby/go-utils/fs"
)

const INPUT_FILE = "input.txt"

func main() {
	input, _ := fs.ReadTextFile(INPUT_FILE)
	total_points := 0
	points := map[byte]int{'A': 1, 'B': 2, 'C': 3}

	for _, ln := range strings.Split(input, "\n") {
		if len(ln) == 0 {
			break
		}

		opponents_move := ln[0]
		my_move := ln[2]

		// expects draw
		if my_move == 'Y' {
			total_points += points[opponents_move] + 3
		}

		// needs to loose
		if my_move == 'X' {
			if opponents_move == 'A' {
				total_points += 3
			}

			if opponents_move == 'B' {
				total_points += 1
			}

			if opponents_move == 'C' {
				total_points += 2
			}
		}

		// needs to win
		if my_move == 'Z' {
			total_points += 6
			if opponents_move == 'A' {
				total_points += 2
			}

			if opponents_move == 'B' {
				total_points += 3
			}

			if opponents_move == 'C' {
				total_points += 1
			}
		}

	}

	fmt.Printf("Total points %d\n", total_points)
}
