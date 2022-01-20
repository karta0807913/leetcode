package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildFromArray(idxMap map[int]int, inorder, postorder []int, offset int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	n := len(postorder) - 1
	leftLength := idxMap[postorder[n]] - offset
	return &TreeNode{
		Val:   postorder[n],
		Left:  buildFromArray(idxMap, inorder[:leftLength], postorder[:leftLength], offset),
		Right: buildFromArray(idxMap, inorder[leftLength+1:], postorder[leftLength:len(postorder)-1], offset+leftLength+1),
	}

}

func buildTree(inorder []int, postorder []int) *TreeNode {
	idxMap := make(map[int]int)
	for idx := range inorder {
		idxMap[inorder[idx]] = idx
	}
	return buildFromArray(idxMap, inorder, postorder, 0)
}
