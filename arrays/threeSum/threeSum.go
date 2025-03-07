package main

import (
	"fmt"
	"sort"
)

// THREE SUM: Given a sorted list of integers, use three pointers, one at the start of the list. The next pointer left is placed right next to the pointer i, and the other pointer right is placed at the end of the list. Move the left and right pointers towards each other and keep checking if the sum of the ints at all the pointers equal the target.
func threeSum(nums []int, target int) ([][]int, error) {
	sort.Ints(nums)
	result := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// Skip duplicates
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result, nil
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	target := 0
	fmt.Println(threeSum(nums, target))
}
