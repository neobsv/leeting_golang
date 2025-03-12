package main

import (
	"fmt"
	"sort"
)

func combination_helper(candidates []int, index int, target int, temp []int, result *[][]int) {

	// fmt.Println("temp: ", temp)

	if index > len(candidates) || target < 0 {
		return
	}

	if target == 0 {
		newres := make([]int, len(temp))
		copy(newres, temp)
		*result = append(*result, newres)
		return
		// fmt.Println("res: ", *result)
	}

	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}

		if target >= candidates[i] {
			temp = append(temp, candidates[i])
			combination_helper(candidates, i+1, target-candidates[i], temp, result)
			temp = temp[:len(temp)-1]
		}
	}

}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	temp := []int{}
	combination_helper(candidates, 0, target, temp, &result)

	// fmt.Println("result: ", result)

	return result
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
