package main

func calLength(s string) (ans int) {
	for i := range s {
		if s[i] == ' ' {
			return i
		}
	}
	return len(s)
}
func lengthOfLastWord(s string) int {
	prev := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			prev = calLength(s[i:])
			i += prev
		}
	}
	return prev
}
