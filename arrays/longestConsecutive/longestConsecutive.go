package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// LONGEST CONSECUTIVE: Make a set of the given integers, and for each "start number" (number whose predecessor is not in the set), calculate the length of the sequence of numbers from start number till the numbers in the set. Maximize this length over all elements.
func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, num := range nums {
		set[num] = true
	}

	res := 0
	for _, num := range nums {
		if set[num-1] {
			continue
		}

		sequence, temp := 1, num+1
		for set[temp] {
			sequence++
			temp++
		}

		res = max(res, sequence)
	}
	return res
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println("Length of longest consecutive elements:", longestConsecutive(nums))

	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println("Length of longest consecutive elements:", longestConsecutive(nums))
}
