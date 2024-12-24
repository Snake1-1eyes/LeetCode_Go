package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str)
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}

	result := make([][]string, 0, len(anagrams))
	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}

func sortString(s string) string {
	splitString := strings.Split(s, "")
	sort.Strings(splitString)
	return strings.Join(splitString, "")
}
