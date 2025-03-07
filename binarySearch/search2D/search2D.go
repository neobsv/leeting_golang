package main

import (
	"fmt"
)

// BINARY SEARCH 2D: In a matrix, convert the 2D indices into 1D indices, row index is division and column index is modulo. Then move through the matrix like we do while performing a binary search on a flat list, and check if the target is found at the middle element.
func binarySearch2D(nums [][]int, target int) ([]int, error) {
	M, N := len(nums), len(nums[0])
	total_elements := M*N - 1
	low, high := 0, total_elements
	result := make([]int, 2)

	for low <= high {
		mid := (low + high) / 2
		mid_row, mid_col := mid/N, mid%N

		// fmt.Println("mid: ", mid_row, mid_col)
		// fmt.Println("mid elem: ", nums[mid_row][mid_col])
		// fmt.Println("Low: ", low, "High: ", high)
		// fmt.Println("target: ", target)

		if nums[mid_row][mid_col] == target {
			result = []int{mid_row, mid_col}
			break
		} else if nums[mid_row][mid_col] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if low > total_elements || high < 0 {
		return nil, fmt.Errorf("Element not found in 2D array")
	}

	return result, nil
}

func main() {
	nums := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	target := 3

	result, _ := binarySearch2D(nums, target)
	fmt.Printf("Element found at position (%d, %d)\n", result[0], result[1])
}
