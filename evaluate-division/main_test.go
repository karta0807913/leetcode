package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type node struct {
	Parent  string
	Product float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	unionFind := make(map[string]*node)
	var find func(current string) *node
	find = func(current string) *node {
		n := unionFind[current]
		if n == nil {
			unionFind[current] = &node{Parent: current, Product: 1}
			return unionFind[current]
		}
		if n.Parent == current {
			return n
		}
		parent := find(n.Parent)
		n.Product *= parent.Product
		n.Parent = parent.Parent
		return parent
	}

	for i, equation := range equations {
		// equation[0] = values[i] * equation[1]
		leftParent := find(equation[0])
		rightParent := find(equation[1])
		rightParent.Parent = leftParent.Parent
		rightParent.Product *= values[i]
		fmt.Printf("unionFind: %v\n", unionFind)
	}

	ans := make([]float64, len(queries))
	for i, query := range queries {
		left := find(query[0])
		right := find(query[1])
		if left.Parent != right.Parent {
			ans[i] = -1
			continue
		}
		ans[i] = unionFind[query[0]].Product / unionFind[query[1]].Product
	}
	return ans
}

func TestEquation(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		[]float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000},
		calcEquation(
			[][]string{
				{"a", "b"}, {"b", "c"},
			},
			[]float64{2, 3},
			[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
		),
	)
}
