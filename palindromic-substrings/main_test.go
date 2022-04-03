package main

func check(s string, left, right int) (ans int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
		ans += 1
	}
	return
}

func countSubstrings(s string) (ans int) {
	for i := range s {
		ans += check(s, i, i)
		ans += check(s, i, i+1)
	}
	return
}
