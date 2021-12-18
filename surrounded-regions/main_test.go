package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func markBoard(x, y int, board [][]byte) bool {
	if x < 0 || len(board[0]) == x || y < 0 || len(board) == y {
		return true
	}
	if board[y][x] == 'X' {
		return false
	} else if board[y][x] == 'O' {
		board[y][x] = 'C'
		a := markBoard(x+1, y, board)
		b := markBoard(x-1, y, board)
		c := markBoard(x, y+1, board)
		d := markBoard(x, y-1, board)
		if a || b || c || d {
			return true
		}
	}
	return false
}

func solve(board [][]byte) {
	for x := 0; x < len(board[0]); x++ {
		markBoard(x, 0, board)
		markBoard(x, len(board)-1, board)
	}

	for y := 0; y < len(board); y++ {
		markBoard(0, y, board)
		markBoard(len(board[0])-1, y, board)
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if board[y][x] == 'C' {
				board[y][x] = 'O'
			} else if board[y][x] == 'O' {
				board[y][x] = 'X'
			}
		}
	}
}

func TestSolve(t *testing.T) {
	var board [][]byte
	board = [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	assert.Equal(t, [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'},
	}, board, nil)

	board = [][]byte{
		{'X', 'O', 'X'},
		{'O', 'X', 'O'},
		{'X', 'O', 'X'},
	}

	solve(board)
	assert.Equal(t, [][]byte{
		{'X', 'O', 'X'},
		{'O', 'X', 'O'},
		{'X', 'O', 'X'},
	}, board, nil)
}
