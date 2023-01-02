package main

import (
	"container/heap"
	"math"
)

type SmallerHeap []int

func (h SmallerHeap) Len() int            { return len(h) }
func (h SmallerHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h SmallerHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h SmallerHeap) Peek() int           { return h[0] }
func (h *SmallerHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *SmallerHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type LargetHeap struct {
	SmallerHeap
}

func (h LargetHeap) Less(i, j int) bool { return h.SmallerHeap[i] > h.SmallerHeap[j] }

type MedianFinder struct {
	lHeap LargetHeap
	rHeap SmallerHeap
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if this.rHeap.Len() == 0 && this.lHeap.Len() == 0 {
		this.lHeap.Push(num)
		return
	}
	if this.FindMedian() < float64(num) {
		heap.Push(&this.rHeap, num)
		if this.lHeap.Len() < this.rHeap.Len() {
			heap.Push(&this.lHeap, heap.Pop(&this.rHeap))
		}
	} else {
		heap.Push(&this.lHeap, num)
		if this.lHeap.Len()-1 > this.rHeap.Len() {
			heap.Push(&this.rHeap, heap.Pop(&this.lHeap))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.rHeap.Len() == 0 && this.lHeap.Len() == 0 {
		return math.MaxFloat64
	}
	if this.rHeap.Len() == 0 {
		return float64(this.lHeap.Peek())
	}

	if this.lHeap.Len() == this.rHeap.Len() {
		return float64(this.lHeap.Peek()+this.rHeap.Peek()) / 2
	}
	return float64(this.lHeap.Peek())
}
