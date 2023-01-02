package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(root *TreeNode, resultMap map[int]int, depth, maxDepth int) int {
	if root == nil {
		return maxDepth
	}
	resultMap[root.Val] = max(resultMap[root.Val], maxDepth)
	maxDepth = max(depth, maxDepth)
	maxDepth = max(dfs(root.Left, resultMap, depth+1, maxDepth), maxDepth)
	maxDepth = max(dfs(root.Right, resultMap, depth+1, maxDepth), maxDepth)
	root.Left, root.Right = root.Right, root.Left
	return maxDepth
}

func treeQueries(root *TreeNode, queries []int) []int {
	result := make([]int, len(queries))
	resultMap := make(map[int]int)
	dfs(root, resultMap, 0, 0)
	dfs(root, resultMap, 0, 0)
	for i, query := range queries {
		result[i] = resultMap[query]
	}
	return result
}
