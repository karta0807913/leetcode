package main

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Cell struct {
	x, y        int
	possibility map[byte]bool
}

type CellHeap []*Cell

func (h CellHeap) Len() int            { return len(h) }
func (h CellHeap) Less(i, j int) bool  { return len(h[i].possibility) < len(h[j].possibility) }
func (h CellHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *CellHeap) Push(x interface{}) { (*h) = append(*h, x.(*Cell)) }
func (h *CellHeap) Pop() interface{} {
	n := len(*h) - 1
	t := (*h)[n]
	*h = (*h)[:n]
	return t
}

type Board struct {
	board       [][]byte
	cols        [9]map[byte]bool
	rows        [9]map[byte]bool
	squares     [9]map[byte]bool
	missingCell CellHeap
}

func (board Board) canPut(x, y int, v byte) bool {
	i := board.calSquareIndex(x, y)
	// if number was exists, return false
	return !(board.cols[y][v] || board.rows[x][v] || board.squares[i][v])
}

func (board Board) CanSolve(x, y int) (byte, bool) {
	b := byte(0)
	for num := byte('1'); num <= byte('9'); num++ {
		if board.canPut(x, y, num) {
			if b != 0 {
				return b, false
			}
			b = num
		}
	}
	return b, true
}

func (board Board) Put(x, y int, b byte) {
	i := board.calSquareIndex(x, y)
	board.cols[y][b] = true
	board.rows[x][b] = true
	board.squares[i][b] = true
	board.board[y][x] = b
	for idx, cell := range board.missingCell {
		if cell.x == x || cell.y == y || board.calSquareIndex(cell.x, cell.y) == i {
			// fmt.Printf("cell: %v\n", cell)
			delete(cell.possibility, b)
			heap.Fix(&board.missingCell, idx)
		}
	}
}

func (board Board) Take(x, y int, b byte) {
	i := board.calSquareIndex(x, y)
	board.cols[y][b] = false
	board.rows[x][b] = false
	board.squares[i][b] = false
	board.board[y][x] = '.'
	for idx := len(board.missingCell) - 1; idx >= 0; idx-- {
		cell := board.missingCell[idx]
		if cell.x == x || cell.y == y || board.calSquareIndex(cell.x, cell.y) == i {
			// fmt.Printf("cell: %v\n", cell)
			if board.canPut(cell.x, cell.y, b) {
				cell.possibility[b] = true
				heap.Fix(&board.missingCell, idx)
			}
		}
	}
}

func (board Board) calSquareIndex(x, y int) int {
	return (x / 3) + ((y / 3) * 3)
}

func NewBoard(board [][]byte) Board {
	b := Board{}
	for i := range b.cols {
		b.cols[i] = make(map[byte]bool)
		b.rows[i] = make(map[byte]bool)
		b.squares[i] = make(map[byte]bool)
	}
	for colIdx, col := range board {
		for rowIdx, val := range col {
			if val == '.' {
				continue
			} else {
				b.cols[colIdx][val] = true
				b.rows[rowIdx][val] = true
				b.squares[b.calSquareIndex(rowIdx, colIdx)][val] = true
			}
		}
	}
	b.board = board
	for colIdx, col := range board {
		for rowIdx, val := range col {
			if val == '.' {
				cell := &Cell{
					x:           rowIdx,
					y:           colIdx,
					possibility: make(map[byte]bool),
				}
				for num := byte('1'); num <= '9'; num++ {
					if b.canPut(rowIdx, colIdx, num) {
						cell.possibility[num] = true
					}
				}
				heap.Push(&b.missingCell, cell)
			}
		}
	}
	return b
}

func recursive(board Board) bool {
	if len(board.missingCell) == 0 {
		return true
	}
	// fmt.Println(board.missingCell[0])
	if len(board.missingCell[0].possibility) == 0 {
		return false
	}
	cell := heap.Pop(&board.missingCell).(*Cell)
	for num := range cell.possibility {
		board.Put(cell.x, cell.y, num)
		// fmt.Printf("put x: %v, y: %v, %c\n", cell.x, cell.y, num)
		if recursive(board) {
			return true
		}
		// fmt.Printf("take x: %v, y: %v, %c\n", cell.x, cell.y, num)
		board.Take(cell.x, cell.y, num)
	}
	heap.Push(&board.missingCell, cell)
	return false
}

func solveSudoku(board [][]byte) {
	// b := NewBoard(board)
	// fmt.Println(heap.Pop(&b.missingCell))
	// for _, cell := range b.missingCell {
	// 	fmt.Printf("cell: %v\n", cell)
	// }
	recursive(NewBoard(board))
}

func TestSolve(t *testing.T) {
	sudoku := [][]byte{
		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
	}
	solveSudoku(sudoku)
	assert.Equal(t, [][]byte{
		{'5', '1', '9', '7', '4', '8', '6', '3', '2'},
		{'7', '8', '3', '6', '5', '2', '4', '1', '9'},
		{'4', '2', '6', '1', '3', '9', '8', '7', '5'},
		{'3', '5', '7', '9', '8', '6', '2', '4', '1'},
		{'2', '6', '4', '3', '1', '7', '5', '9', '8'},
		{'1', '9', '8', '5', '2', '4', '3', '6', '7'},
		{'9', '7', '5', '8', '6', '3', '1', '2', '4'},
		{'8', '3', '2', '4', '9', '1', '7', '5', '6'},
		{'6', '4', '1', '2', '7', '5', '9', '8', '3'},
	}, sudoku)

	sudoku = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(sudoku)
	assert.Equal(t, [][]byte{
		{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
		{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
		{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
		{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
		{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
		{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
		{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
		{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
		{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
	}, sudoku)
}
