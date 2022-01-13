package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func _exist(board [][]byte, x, y int, word string) bool {
	if len(word) == 0 {
		return true
	}
	if board[y][x] != word[0] {
		return false
	}
	c := board[y][x]
	board[y][x] = 0
	result := false
	word = word[1:]
	if len(word) == 0 {
		return true
	}
	if x+1 != len(board[0]) {
		result = result || _exist(board, x+1, y, word)
	}
	if x != 0 {
		result = result || _exist(board, x-1, y, word)
	}

	if y+1 != len(board) {
		result = result || _exist(board, x, y+1, word)
	}

	if y != 0 {
		result = result || _exist(board, x, y-1, word)
	}

	board[y][x] = c
	return result
}

func exist(board [][]byte, word string) bool {
	if len(board)*len(board[0]) < len(word) {
		return false
	}
	for y, column := range board {
		for x := range column {
			if _exist(board, x, y, word) {
				return true
			}
		}
	}
	return false
}

func TestFunc(t *testing.T) {
	assert.True(t, exist([][]byte{[]byte("a")}, "a"))
}
