package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

type ListMap map[*Node]*Node

func (lm ListMap) searchOrCreate(addr *Node) *Node {
	if addr == nil {
		return nil
	}
	res, ok := lm[addr]
	if !ok {
		res = &Node{
			Val: addr.Val,
		}
	}
	return res
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	lm := make(ListMap)
	newHead := new(Node)
	runner := &newHead
	for head != nil {
		*runner = lm.searchOrCreate(head)
		(*runner).Random = lm.searchOrCreate(head.Random)
		runner = &(*runner).Next
		head = head.Next
	}
	return newHead
}
