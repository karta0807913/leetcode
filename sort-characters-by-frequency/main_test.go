package main

import "strings"

func frequencySort(s string) string {
	bucket := make([][]byte, len(s)+1)
	frequency := make(map[byte]int, len(s))
	for i := range s {
		frequency[s[i]] += 1
	}
	for k, v := range frequency {
		bucket[v] = append(bucket[v], k)
	}
	result := strings.Builder{}
	for n := len(bucket) - 1; n >= 0; n-- {
		for _, char := range bucket[n] {
			for i := 0; i < n; i++ {
				result.WriteByte(char)
			}
		}
	}
	return result.String()
}
