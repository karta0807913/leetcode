package main

func isSubsequence(s string, t string) bool {
	si, ti := 0, 0
	for ; si != len(s) && ti != len(t); ti++ {
		if s[si] == t[ti] {
			si++
		}
	}
	return si == len(s)
}
