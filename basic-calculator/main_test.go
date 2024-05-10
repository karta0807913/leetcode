package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Operator int

const (
	Plus Operator = 1
	Sub  Operator = 2
)

func readDigits(s string) (int, string) {
	for s[0] == ' ' {
		s = s[1:]
	}
	base := 1
	if s[0] == '-' {
		base = -1
		s = s[1:]
	}
	for s[0] == ' ' {
		s = s[1:]
	}
	if s[0] == '(' {
		num, s := calculateSubExpression(s[1:])
		return base * num, s
	}
	num := 0
	for len(s) != 0 {
		if s[0] == ' ' {
			s = s[1:]
			continue
		}
		if s[0] == '+' || s[0] == '-' || s[0] == ')' {
			return num * base, s
		}
		num *= 10
		num += int(s[0] - '0')
		s = s[1:]
	}
	return num * base, s
}

func readOperator(s string) (Operator, string) {
	for s[0] == ' ' {
		s = s[1:]
	}
	switch s[0] {
	case '-':
		return Sub, s[1:]
	case '+':
		return Plus, s[1:]
	}
	return 0, s[1:]
}

func calculateSubExpression(s string) (int, string) {
	var operator Operator
	var digit, result int
	result, s = readDigits(s)
	for len(s) != 0 {
		if s[0] == ' ' {
			s = s[1:]
			continue
		}
		if s[0] == ')' {
			return result, s[1:]
		}
		operator, s = readOperator(s)
		digit, s = readDigits(s)
		switch operator {
		case Plus:
			result += digit
		case Sub:
			result -= digit
		}
	}
	return result, s
}

func calculate(s string) int {
	result, _ := calculateSubExpression(s)
	return result
}

func TestCalculator(t *testing.T) {
	assert := require.New(t)
	assert.Equal(-12, calculate("- (3 + (4 + 5))"))
	assert.Equal(-12, calculate("- (3 + (4 + 5))"))
	assert.Equal(3, calculate("   (  3 ) "))
	assert.Equal(3, calculate("3"))
	assert.Equal(6, calculate("1+3+4-2"))
	assert.Equal(-6, calculate("-(1+3+4-2)"))
	assert.Equal(-4, calculate("-((1+3+4-2)+ -2)"))
	assert.Equal(-2, calculate("1+ 2-3+2-((1+3+4-2)+ -2)"))
	assert.Equal(-2, calculate("(-2)"))
	assert.Equal(-1, calculate("1+  (-2)"))
}
