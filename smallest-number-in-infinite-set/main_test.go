package main

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SmallestInfiniteSet struct {
	heap      IntHeap
	nonExists map[int]bool
	smallest  int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		heap:      IntHeap{},
		nonExists: make(map[int]bool),
		smallest:  1,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if len(this.heap) == 0 {
		this.nonExists[this.smallest] = true
		n := this.smallest
		this.smallest += 1
		return n
	}
	n := heap.Pop(&this.heap).(int)
	this.nonExists[n] = true

	return n
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if !this.nonExists[num] {
		return
	}
	this.nonExists[num] = false
	heap.Push(&this.heap, num)
}
