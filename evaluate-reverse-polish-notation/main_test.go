package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0, 10)
	for _, token := range tokens {
		switch token {
		case "+":
			stack[len(stack)-2] = stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] = stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] = stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] = stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			i, _ := strconv.Atoi(token)
			stack = append(stack, i)
		}
	}
	return stack[0]
}
