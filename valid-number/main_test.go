package main

func isInteger(s string) bool {
	if hasPrefix(s) {
		s = s[1:]
	}
	if len(s) == 0 {
		return false
	}
	for i := range s {
		if s[i] < '0' || '9' < s[i] {
			return false
		}
	}
	return true
}

func hasPrefix(s string) bool {
	return len(s) > 0 && (s[0] == '-' || s[0] == '+')
}

func isNumber(s string) bool {
	if hasPrefix(s) {
		s = s[1:]
		if len(s) == 0 {
			return false
		}
	}
	hasDot := false
	valid := false
	for len(s) != 0 {
		if s[0] == '.' {
			if hasDot {
				return false
			}
			hasDot = true
		} else if s[0] == 'E' || s[0] == 'e' {
			if !valid {
				return false
			}
			return isInteger(s[1:])
		} else if s[0] < '0' || '9' < s[0] {
			return false
		} else {
			valid = true
		}
		s = s[1:]
	}
	return valid
}
