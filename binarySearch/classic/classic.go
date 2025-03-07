package main

import (
	"fmt"
	"sort"
)

// BINARY SEARCH: Just a search to check for the presence of an element in a list, if it exists return an index, otherwise return -1.
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// FIND: A variant of binarySearch which returns the earliest index of the target element in a sorted list which contians duplicate elements, return -1 if it doesn't exist
func find(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			for target == arr[mid] {
				mid--
			}
			if mid < 0 {
				return 0
			} else {
				return mid + 1
			}
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 11, 11, 11, 11, 11, 11, 13, 15, 17, 19}
	sort.Ints(arr)
	target := 11

	index := binarySearch(arr, target)
	fmt.Println("Sorted array, target, index of element:", arr, target, index)

	find_index := find(arr, target)
	fmt.Println("Sorted array, target, index of element:", arr, target, find_index)
}
