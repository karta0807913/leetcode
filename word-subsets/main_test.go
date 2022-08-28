package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func wordSubsets(words1 []string, words2 []string) []string {
	result := []string{}
	general := make(map[byte]int, 26)
	for _, word := range words2 {
		tmp := make(map[byte]int, 26)
		for i := range word {
			tmp[word[i]] += 1
		}
		for key, val := range tmp {
			general[key] = max(general[key], val)
		}
	}
C:
	for _, word := range words1 {
		tmp := make(map[byte]int, 26)
		for i := range word {
			tmp[word[i]] += 1
		}
		for key, val := range general {
			if val > tmp[key] {
				continue C
			}
		}
		result = append(result, word)
	}
	return result
}
