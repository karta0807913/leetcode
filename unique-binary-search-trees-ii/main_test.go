package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func cloneNode(beginVal int, node *TreeNode) (int, *TreeNode) {
	if node == nil {
		return beginVal, nil
	}
	count, clonedNode := cloneNode(beginVal, node.Left)
	result := &TreeNode{Val: count + 1, Left: clonedNode}
	count, clonedNode = cloneNode(count+1, node.Right)
	result.Right = clonedNode
	return count, result
}

func generateTrees(n int) []*TreeNode {
	cache := make([][]*TreeNode, n+1)
	cache[0] = []*TreeNode{}
	cache[1] = []*TreeNode{{Val: 0}}
	for i := 2; i <= n; i++ {
		for left := 0; left < i; left++ {
			right := i - left - 1
			if left == 0 {
				for rightNodes := range cache[right] {
					_, right := cloneNode(1, cache[right][rightNodes])
					cache[i] = append(cache[i], &TreeNode{Val: 1, Right: right})
				}
			} else if right == 0 {
				for leftNodes := range cache[left] {
					val, left := cloneNode(0, cache[left][leftNodes])
					fmt.Printf("val: %v\n", val)
					cache[i] = append(cache[i], &TreeNode{Val: val + 1, Left: left})
				}
			} else {
				for leftNodes := range cache[left] {
					for rightNodes := range cache[right] {
						val, left := cloneNode(0, cache[left][leftNodes])
						_, right := cloneNode(val+1, cache[right][rightNodes])
						cache[i] = append(cache[i], &TreeNode{Val: val + 1, Left: left, Right: right})
					}
				}
			}
		}
	}
	return cache[n]
}

func TestTrees(t *testing.T) {
	assert.ElementsMatch(t, []*TreeNode{{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}, {
		Val:  2,
		Left: &TreeNode{Val: 1},
	}}, generateTrees(2))
}
