package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tlaceby/2022AOC/lib"
)

func main() {
	const TOP_COUNT = 3
	lines := lib.GetLines("input.txt")
	top_callories := make([]int, TOP_COUNT)
	current_callories := 0

	for _, line := range lines {
		// If the length is one then is simply means a newline
		if len(line) == 1 {
			// Check the top elfs and see if elfs is greater than any of them
			for indx, cal := range top_callories {
				// If the number is larger than top N then swap the value at the lowest index
				if current_callories > cal {
					top_callories[indx] = current_callories
					break
				}
			}

			// Resets the elf for the next iteration through
			current_callories = 0
		}

		if len(line) > 1 {
			// remove the trailing \r character on the line
			line = strings.ReplaceAll(line, "\r", "")
			num, _ := strconv.Atoi(line)

			// running sum of callories
			current_callories += num
		}
	}

	fmt.Printf("Sum %d\n", lib.Accumulate(top_callories))
}
