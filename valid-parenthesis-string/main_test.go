package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func checkValidString(s string) bool {
// 	left, golden := 0, 0
// 	for _, val := range s {
// 		if val == '(' {
// 			left += 1
// 		} else if val == ')' {
// 			left -= 1
// 		} else if val == '*' {
// 			golden += 1
// 		}
// 		if left != 0 && golden != 0 {
// 			golden -= 1
// 			if left < 0 {
// 				left += 1
// 			} else {
// 				left -= 1
// 			}
// 		}
// 	}
// 	fmt.Println(left, golden)
// 	return left == 0
// }

func checkValidString(s string) bool {
	lo, hi := 0, 0
	for _, val := range s {
		if val == '(' {
			lo += 1
		} else {
			lo -= 1
		}
		if val != ')' {
			hi += 1
		} else {
			hi -= 1
		}
		if hi < 0 {
			break
		}
		if lo < 0 {
			lo = 0
		}
	}
	return lo == 0
}

func TestString(t *testing.T) {
	assert.False(t, checkValidString("())*"), nil)
	assert.True(t, checkValidString("()"), nil)
	assert.True(t, checkValidString("(*)"), nil)
	assert.True(t, checkValidString("(*))"), nil)
	assert.True(t, checkValidString("(((*))"), nil)
	assert.False(t, checkValidString("((())"), nil)
	assert.True(t, checkValidString("(**)))(*)"), nil)
	assert.False(t, checkValidString("(*)))(*)"), nil)
	assert.False(t, checkValidString("(*)("), nil)
	assert.False(t, checkValidString("(**()("), nil)
	assert.False(t, checkValidString("(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"), nil)
}
