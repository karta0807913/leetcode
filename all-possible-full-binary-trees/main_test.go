package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		return []*TreeNode{}
	}
	cache := make([][]*TreeNode, (n+1)/2)
	cache[0] = []*TreeNode{{Val: 0}}
	target := (n + 1) / 2
	for i := 1; i < target; i++ {
		cache[i] = make([]*TreeNode, 0)
		for t := 0; t < i; t++ {
			for _, leftNode := range cache[t] {
				for _, rightNode := range cache[i-t-1] {
					cache[i] = append(cache[i], &TreeNode{
						Val:   0,
						Left:  leftNode,
						Right: rightNode,
					})
				}
			}
		}
	}
	return cache[len(cache)-1]
}
