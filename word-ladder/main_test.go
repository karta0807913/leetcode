package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := make(map[string]int)
	for i, word := range wordList {
		dict[word] = i
	}
	if _, ok := dict[endWord]; !ok {
		return 0
	}
	if _, ok := dict[beginWord]; !ok {
		wordList = append(wordList, beginWord)
		dict[beginWord] = len(wordList) - 1
	}
	queue := make([]int, 1, len(dict))
	queue[0] = dict[beginWord]
	next := make([]int, 0, len(dict))
	qh := queue
	nh := next
	num := 1
	for len(queue) != 0 {
		num++
		for _, wordIdx := range queue {
			word := wordList[wordIdx]
			for idx := range word {
				for i := 0; i < 26; i++ {
					nextWord := word[:idx] + string(rune('a'+i)) + word[idx+1:]
					if idx, ok := dict[nextWord]; ok {
						if nextWord == endWord {
							return num
						}
						delete(dict, nextWord)
						next = append(next, idx)
					}
				}
			}
		}
		queue = next
		next = qh[:0]
		qh, nh = nh, qh
	}
	return 0
}
