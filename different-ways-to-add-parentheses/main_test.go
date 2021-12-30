package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

type NumberCell struct {
	Val      int
	Operator func(int, int) int
}

func decodeExpression(expression string) []NumberCell {
	cells := make([]NumberCell, 0)
	currentCell := NumberCell{}
	for _, val := range expression {
		if '0' <= val && val <= '9' {
			currentCell.Val = 10*currentCell.Val + int(val-'0')
		} else {
			switch val {
			case '*':
				currentCell.Operator = mul
			case '-':
				currentCell.Operator = sub
			case '+':
				currentCell.Operator = add
			}
			cells = append(cells, currentCell)
			currentCell = NumberCell{}
		}
	}
	cells = append(cells, currentCell)
	return cells
}

func cal(cells []NumberCell) []int {
	if len(cells) == 1 {
		return []int{cells[0].Val}
	}
	if len(cells) == 2 {
		return []int{cells[0].Operator(cells[0].Val, cells[1].Val)}
	}
	if len(cells) == 3 {
		return []int{
			cells[1].Operator(
				cells[0].Operator(cells[0].Val, cells[1].Val), cells[2].Val,
			),
			cells[0].Operator(
				cells[0].Val, cells[1].Operator(cells[1].Val, cells[2].Val),
			),
		}
	}
	panic("cell count is not 2 or 3")
}

func recursive(cells []NumberCell, start, end int) []int {
	if end-start < 3 {
		return cal(cells[start : end+1])
	}
	answer := make([]int, 0)
	for idx := start; idx < end; idx++ {
		left := recursive(cells, start, idx)
		right := recursive(cells, idx+1, end)
		for _, leftVal := range left {
			for _, rightVal := range right {
				answer = append(answer, cells[idx].Operator(leftVal, rightVal))
			}
		}
	}
	return answer
}

func diffWaysToCompute(expression string) []int {
	cells := decodeExpression(expression)
	return recursive(cells, 0, len(cells)-1)
}

func TestWays(t *testing.T) {
	// assert.ElementsMatch(t, []int{0, 2}, diffWaysToCompute("2-1-1"))
	assert.ElementsMatch(t, []int{
		1, 3, 1, 1, -1,
	}, diffWaysToCompute("2-1-1-1"))
	assert.ElementsMatch(t, []int{
		-34, -14, -10, -10, 10,
	}, diffWaysToCompute("2*3-4*5"))
}
