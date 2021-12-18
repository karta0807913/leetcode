package main

import (
	"fmt"
	"strconv"
	"testing"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

func (node Node) String() string {
	return strconv.Itoa(node.Val)
}

func iterNodes(node *Node, nodes []*Node) {
	if nodes[node.Val] != nil {
		return
	}
	nodes[node.Val] = node
	for idx := range node.Neighbors {
		iterNodes(node.Neighbors[idx], nodes)
	}
}

func createNewNode() *Node {
	return &Node{Neighbors: make([]*Node, 0)}
}

func cloneNode(oldNode, newNode *Node, newNodes []*Node) {
	for _, neighbor := range oldNode.Neighbors {
		if newNodes[neighbor.Val] == nil {
			newNodes[neighbor.Val] = createNewNode()
		}
		newNode.Neighbors = append(newNode.Neighbors, newNodes[neighbor.Val])
	}
	newNode.Val = oldNode.Val
}

func cloneGraph(node *Node) *Node {
	nodes := make([]*Node, 101)
	iterNodes(node, nodes)
	fmt.Println(nodes[:5])
	for _, n := range nodes[1:5] {
		fmt.Println(n.Neighbors)
	}
	newNodes := make([]*Node, 101)
	for idx := range nodes {
		if nodes[idx] == nil {
			continue
		}
		if newNodes[idx] == nil {
			newNodes[idx] = createNewNode()
		}
		cloneNode(nodes[idx], newNodes[idx], newNodes)
	}
	for _, n := range newNodes {
		if n != nil {
			return n
		}
	}
	return &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
}

func TestGraph(t *testing.T) {
	node4 := &Node{4, []*Node{&Node{}, &Node{}}}
	node3 := &Node{3, []*Node{&Node{}, &Node{}}}
	node2 := &Node{2, []*Node{&Node{}, node3}}
	node1 := &Node{1, []*Node{node2, node4}}
	node2.Neighbors[0] = node1
	node3.Neighbors[0] = node2
	node3.Neighbors[1] = node4
	node4.Neighbors[1] = node4
	node4.Neighbors[0] = node1
	node4.Neighbors[1] = node3
	newNode := cloneGraph(node1)
	fmt.Println(newNode)
}
