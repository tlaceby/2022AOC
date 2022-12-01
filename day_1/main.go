package main

import (
	"fmt"
	"strconv"

	"github.com/tlaceby/go-utils/arrays"
	"github.com/tlaceby/go-utils/fs"
)

func main() {
	const TOP_COUNT = 3
	lines, _ := fs.GetLines("input.txt")
	top_callories := make([]int, TOP_COUNT)
	current_callories := 0

	for _, line := range lines {
		println(len(line))
		// If the length is one then is simply means a newline
		if len(line) == 0 {
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

		if len(line) > 0 {
			num, _ := strconv.Atoi(line)
			current_callories += num
		}
	}

	fmt.Printf("Sum %d %v\n", arrays.Accumulate(top_callories), top_callories)
}
