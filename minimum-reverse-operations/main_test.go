package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

type GridMove struct {
	LeftRequired  int
	RightRequired int
	Step          int
}

func (m *GridMove) CanMove(i, n int) bool {
	return i-m.LeftRequired >= 0 && i+m.RightRequired < n
}

func minReverseOperations(n int, p int, banned []int, k int) []int {
	steps := make([]GridMove, 0, k)
	for i := 0; i < k; i++ {
		steps = append(steps, GridMove{
			LeftRequired:  i,
			Step:          k - i*2 - 1,
			RightRequired: k - i - 1,
		})
	}
	ans := make([]int, n)
	for _, i := range banned {
		ans[i] = -1
	}
	current := []int{p}
	next := []int{}
	stepCount := 1
	for len(current) != 0 {
		for _, idx := range current {
			remainRight := (n - idx - 1)
			startPoint := sort.Search(len(steps), func(i int) bool {
				return steps[i].RightRequired <= remainRight
			})
			for _, step := range steps[startPoint:] {
				nextIdx := idx + step.Step
				if !step.CanMove(idx, n) {
					break
				}
				if 0 <= nextIdx && nextIdx < n && ans[nextIdx] == 0 {
					ans[nextIdx] = stepCount
					next = append(next, nextIdx)
				}
			}
		}
		current, next = next, current
		next = next[:0]
		stepCount++
	}
	for i := range ans {
		if ans[i] == 0 {
			ans[i] = -1
		}
	}
	ans[p] = 0
	return ans
}

func TestStep(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		[]int{-1, -1, 0},
		minReverseOperations(3, 2, []int{}, 1),
	)
	assert.Equal(
		[]int{0, -1, -1, 1},
		minReverseOperations(4, 0, []int{1, 2}, 4),
	)
	assert.Equal(
		[]int{0, -1, -1, -1, -1},
		minReverseOperations(5, 0, []int{2, 4}, 3),
	)
	assert.Equal(
		[]int{-1, -1, 0, -1},
		minReverseOperations(4, 2, []int{0, 1, 3}, 1),
	)
}
