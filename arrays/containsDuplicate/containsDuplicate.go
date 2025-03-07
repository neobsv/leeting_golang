package main

import (
	"fmt"
)

// CONTAINS DUPLICATE: Use a hashmap to check for duplicates in the list of ints.
func containsDuplicate(nums []int) bool {
	hashr := make(map[int]bool)

	for _, num := range nums {
		if _, ok := hashr[num]; ok {
			return true
		}
		hashr[num] = true
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 5}
	fmt.Println(containsDuplicate(nums)) // Output: true
}
