package main

import (
	"fmt"
	"sort"
)

func topKFrequentElements(words []string, k int) []string {
	counter := make(map[string]int, len(words))
	var unique []string

	for _, word := range words {
		if _, found := counter[word]; !found {
			unique = append(unique, word)
		}
		counter[word]++
	}

	sort.SliceStable(unique, func(i, j int) bool {
		if counter[unique[i]] == counter[unique[j]] {
			// frequencies are equal, compare the words
			return unique[i] < unique[j]
		}
		// compare the frequencies
		return counter[unique[i]] > counter[unique[j]]
	})

	return unique[:k]
}

func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	fmt.Println(topKFrequentElements(words, k))
}
