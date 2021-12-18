package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isValid(str string) bool {
	stack := make([]rune, 0)
	for _, s := range str {
		if s == '(' {
			stack = append(stack, '(')
		} else if s == '{' {
			stack = append(stack, '{')
		} else if s == '[' {
			stack = append(stack, '[')
		} else {
			if len(stack) == 0 {
				return false
			}
			if s == '}' {
				if stack[len(stack)-1] != '{' {
					return false
				}
			} else if s == ']' {
				if stack[len(stack)-1] != '[' {
					return false
				}
			} else if s == ')' {
				if stack[len(stack)-1] != '(' {
					return false
				}
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	assert.Equal(t, true, isValid("()"), nil)
	assert.Equal(t, true, isValid("()[]{}"), nil)
	assert.Equal(t, true, isValid("{[]}"), nil)
	assert.Equal(t, false, isValid("([)]"), nil)
	assert.Equal(t, false, isValid("]"), nil)
}
