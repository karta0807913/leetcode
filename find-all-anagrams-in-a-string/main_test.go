package main

func findAnagrams(s string, p string) (ans []int) {
	if len(s) < len(p) {
		return []int{}
	}
	m := make(map[byte]int)
	current := make(map[byte]int)
	for i := range p {
		m[p[i]] += 1
	}
	c := 0
	for idx := 0; idx < len(p); idx++ {
		s := s[idx]
		current[s] += 1
		if m[s] != 0 && current[s] == m[s] {
			c += 1
		}
	}
	if c == len(m) {
		ans = append(ans, 0)
	}
	for idx := len(p); idx < len(s); idx++ {
		n := idx - len(p)
		if current[s[n]] == m[s[n]] {
			c -= 1
		}
		current[s[n]] -= 1
		current[s[idx]] += 1
		if m[s[idx]] != 0 && current[s[idx]] == m[s[idx]] {
			c += 1
		}
		if c == len(m) {
			ans = append(ans, n)
		}
	}
	return
}
