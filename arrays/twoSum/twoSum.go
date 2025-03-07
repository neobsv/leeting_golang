package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	hashr := make(map[int]int, len(nums))
	for i, num := range nums {
		if other, found := hashr[num]; found {
			return []int{other, i}
		}
		hashr[target-num] = i
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
