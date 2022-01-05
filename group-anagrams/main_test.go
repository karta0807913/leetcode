package main

import (
	"sort"
)

type Strings struct {
	data []byte
}

func (s *Strings) Len() int {
	return len(s.data)
}

func (s Strings) Less(i, j int) bool {
	return s.data[i] < s.data[j]
}

func (s *Strings) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func groupAnagrams(strs []string) (ans [][]string) {
	m := make(map[string][]string)
	for _, key := range strs {
		result := &Strings{data: []byte(key)}
		sort.Sort(result)
		m[string(result.data)] = append(m[string(result.data)], key)
	}
	for _, val := range m {
		ans = append(ans, val)
	}
	return ans
}
