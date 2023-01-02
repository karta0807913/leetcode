package main

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func iterate(root *TreeNode, queries [][2]int, result [][]int, idx int) int {
	if root == nil || idx == len(queries) {
		return idx
	}
	idx = iterate(root.Left, queries, result, idx)
	if root.Val == queries[idx][1] {
		result[idx][1] = root.Val
	}
	for root.Val > queries[idx][1] {
		if idx != 0 && result[idx][0] == -1 {
			result[idx][0] = result[idx-1][0]
		}
		result[idx][1] = root.Val
		idx++
		if idx == len(queries) {
			return idx
		}
	}
	result[idx][0] = root.Val
	idx = iterate(root.Right, queries, result, idx)
	return idx
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	result := make([][]int, 0, len(queries))
	sortedQueries := make([][2]int, 0, len(queries))
	for i, query := range queries {
		result = append(result, []int{-1, -1})
		sortedQueries = append(sortedQueries, [2]int{i, query})
	}
	sort.Slice(sortedQueries, func(i, j int) bool {
		return sortedQueries[i][1] < sortedQueries[j][1]
	})
	idx := iterate(root, sortedQueries, result, 0)
	for ; idx < len(queries)-1; idx++ {
		result[idx+1][0] = result[idx][0]
	}
	finalResult := make([][]int, len(queries))
	for i := range sortedQueries {
		finalResult[sortedQueries[i][0]] = result[i]
	}
	return finalResult
}
