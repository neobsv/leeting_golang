package main

import (
	"sort"
)

// SORT STRING: Convert the string into a slice of runes and sort it using the sort.Slice function. Sorting a slice of runes will arrange them in ascending order.
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

// GROUP ANAGRAMS: Parse the list of words and make a map of string keyed lists, and then sort the each individual word. If word a and word b are anagrams, then sorting them will return the same sequence. Use this sequence as a key for the hashmap and add all the similar words to the list in the hashmap.
func groupAnagrams(words []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, word := range words {
		sortedWord := sortString(word)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func main() {
	words := []string{"listen", "silent", "potato", "stop", "tops", "taps", "pots", "sot"}
	groups := groupAnagrams(words)

	for _, group := range groups {
		println(group)
	}
	// Output: [listen silent] [potato stop] [tops taps] [pots sot]
}
