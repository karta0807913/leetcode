package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

type Node struct {
	Parent int
	Score  int
}

func FindParent(node int, parentMap map[int]*Node) *Node {
	if parentMap[node] == nil {
		parentMap[node] = &Node{
			Parent: node,
			Score:  math.MaxInt,
		}
		return parentMap[node]
	}
	parentNode := parentMap[node]
	minScore := parentNode.Score
	findedArray := []int{}
	for parentNode.Parent != node {
		findedArray = append(findedArray, node)
		node = parentNode.Parent
		parentNode = parentMap[node]
		if minScore > parentNode.Score {
			minScore = parentNode.Score
		}
	}
	for _, n := range findedArray {
		parentMap[n] = parentNode
	}
	parentNode.Score = minScore
	return parentNode
}

func SetParent(parent, node int, parentMap map[int]*Node) int {
	parentNode := FindParent(parent, parentMap)
	currentNode := FindParent(node, parentMap)
	if parentNode.Score > currentNode.Score {
		parentNode.Score = currentNode.Score
	}
	parentMap[currentNode.Parent] = parentNode
	return parentNode.Parent
}

func minScore(n int, roads [][]int) int {
	parentMap := make(map[int]*Node)
	for _, road := range roads {
		parent := FindParent(road[0], parentMap)
		SetParent(parent.Parent, road[1], parentMap)
		if parentMap[parent.Parent].Score > road[2] {
			parentMap[parent.Parent].Score = road[2]
		}
	}
	node := FindParent(n, parentMap)
	return node.Score
}

func TestScore(t *testing.T) {
	assert := require.New(t)
	assert.Equal(124, minScore(79, [][]int{
		{16, 42, 3629}, {50, 60, 9234}, {37, 43, 5521}, {24, 75, 6423}, {52, 24, 2093}, {1, 39, 8856}, {55, 49, 7632}, {49, 69, 7593}, {56, 14, 2513}, {2, 31, 3143}, {24, 46, 1016}, {47, 38, 6891}, {43, 48, 3497}, {59, 67, 4091}, {48, 11, 7615}, {2, 65, 7389}, {34, 59, 6375}, {51, 76, 5172}, {52, 47, 6546}, {51, 2, 3899}, {61, 50, 732}, {19, 52, 3133}, {56, 51, 7253}, {64, 75, 8574}, {60, 2, 2866}, {76, 12, 9550}, {14, 79, 7345}, {75, 45, 8175}, {9, 78, 5951}, {42, 61, 4694}, {74, 26, 4640}, {65, 3, 3635}, {11, 74, 5242}, {12, 56, 6915}, {46, 55, 4106}, {38, 54, 4533}, {69, 37, 1557}, {78, 16, 4688}, {3, 46, 9299}, {26, 44, 3181}, {44, 76, 3744}, {24, 42, 2049}, {11, 24, 5557}, {9, 24, 124}, {45, 34, 2511}, {31, 1, 9796}, {54, 64, 5522}, {69, 51, 5384}, {39, 19, 2828}, {67, 9, 7133},
	}))
	assert.Equal(2, minScore(4, [][]int{
		{1, 2, 2}, {1, 3, 4}, {3, 4, 7},
	}))
}
