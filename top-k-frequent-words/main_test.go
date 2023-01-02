package main

import "sort"

func topKFrequent(words []string, k int) []string {
	wordCount := make(map[string]int)
	for _, w := range words {
		wordCount[w] += 1
	}
	words = words[:0]
	for w := range wordCount {
		words = append(words, w)
	}
	sort.Slice(words, func(i, j int) bool {
		sI, sJ := words[i], words[j]
		if wordCount[sI] > wordCount[sJ] {
			return true
		}
		if wordCount[sI] == wordCount[sJ] {
			return sI < sJ
		}
		return false
	})
	return words[:k]
}
