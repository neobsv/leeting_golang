package main

import (
	"fmt"
	"sort"
)

func combination_helper(candidates []int, index int, target int, temp []int, result *[][]int) {

	// fmt.Println("temp: ", temp)

	if target == 0 {
		newres := make([]int, len(temp))
		copy(newres, temp)
		*result = append(*result, newres)
		return
		// fmt.Println("res: ", *result)
	}

	for i := index; i < len(candidates); i++ {
		if target >= candidates[i] {
			temp = append(temp, candidates[i])
			combination_helper(candidates, i, target-candidates[i], temp, result)
			temp = temp[:len(temp)-1]
		}
	}

}

func combinationSum(candidates []int, target int) [][]int {

	sort.Ints(candidates)
	result := [][]int{}
	temp := []int{}
	combination_helper(candidates, 0, target, temp, &result)

	fmt.Println("result: ", result)
	return result
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	combinationSum(candidates, target)
}
