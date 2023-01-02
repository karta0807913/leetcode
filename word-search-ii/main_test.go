package main

type TieNode struct {
	Word     string
	Children map[byte]*TieNode
}

func recursive(board [][]byte, x, y int, tieNode *TieNode, result []string) []string {
	if y < 0 || len(board) <= y {
		return result
	}
	if x < 0 || len(board[y]) <= x {
		return result
	}
	tieNode = tieNode.Children[board[y][x]]
	if tieNode == nil {
		return result
	}
	if tieNode.Word != "" {
		result = append(result, tieNode.Word)
		tieNode.Word = ""
	}

	tmp := board[y][x]
	board[y][x] = 0
	result = recursive(board, x+1, y, tieNode, result)
	result = recursive(board, x-1, y, tieNode, result)
	result = recursive(board, x, y+1, tieNode, result)
	result = recursive(board, x, y-1, tieNode, result)
	board[y][x] = tmp
	return result
}

func findWords(board [][]byte, words []string) []string {
	tieHead := TieNode{Children: make(map[byte]*TieNode)}
	for _, word := range words {
		node := &tieHead
		for index := range word {
			if node.Children == nil {
				node.Children = make(map[byte]*TieNode)
			}
			nextNode := node.Children[word[index]]
			if nextNode == nil {
				nextNode = &TieNode{}
				node.Children[word[index]] = nextNode
			}
			node = nextNode
		}
		node.Word = word
	}
	var result []string
	for y, column := range board {
		for x := range column {
			result = recursive(board, x, y, &tieHead, result)
		}
	}
	return result
}
