package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TieNode struct {
	char   byte
	isLeaf bool
	Next   []*TieNode
}

type SearchState struct {
	currentNode *TieNode
	idx         int
}

type WordDictionary struct {
	head *TieNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		head: &TieNode{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	pointer := this.head
	for idx := range word {
		var next *TieNode
		for _, node := range pointer.Next {
			if node.char == word[idx] {
				next = node
				break
			}
		}
		if next == nil {
			next = &TieNode{
				char: word[idx],
			}
			pointer.Next = append(pointer.Next, next)
		}
		pointer = next
	}
	pointer.isLeaf = true
}

func (this *WordDictionary) Search(word string) bool {
	stack := make([]SearchState, 1, len(word))
	stack[0].idx = 0
	stack[0].currentNode = this.head
	for len(stack) != 0 {
		n := len(stack) - 1
		search := stack[n]
		stack = stack[:n]
		if search.idx == len(word) {
			if search.currentNode.isLeaf {
				return true
			}
			continue
		}
		for _, node := range search.currentNode.Next {
			if node.char == word[search.idx] || word[search.idx] == '.' {
				stack = append(stack, SearchState{
					idx:         search.idx + 1,
					currentNode: node,
				})
			}
		}
	}
	return false
}

func TestWordTie(t *testing.T) {
	dic := Constructor()
	dic.AddWord("abc")
	assert.True(t, dic.Search("..."))
	dic = Constructor()
	dic.AddWord("a")
	assert.False(t, dic.Search("..."))
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
