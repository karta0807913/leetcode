package main

import (
	"fmt"
	"testing"
)

type LinkedList struct {
	Key  int
	Val  int
	Next *LinkedList
}

type LRUCache struct {
	hashMap  map[int]*LinkedList
	listHead *LinkedList
	listTail *LinkedList
	capacity int
}

func Constructor(capacity int) LRUCache {
	head := &LinkedList{}
	return LRUCache{
		hashMap:  make(map[int]*LinkedList, capacity),
		listHead: head,
		listTail: head,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	prevNode := this.hashMap[key]
	if prevNode == nil {
		return -1
	}
	if prevNode == this.listHead {
		return prevNode.Next.Val
	}
	node := prevNode.Next
	if node == this.listTail {
		this.listTail = prevNode
	}
	prevNode.Next = node.Next
	if prevNode.Next != nil {
		this.hashMap[prevNode.Next.Key] = prevNode
	}
	if this.listHead.Next != nil {
		this.hashMap[this.listHead.Next.Key] = node
	}
	this.hashMap[key] = this.listHead
	node.Next, this.listHead.Next = this.listHead.Next, node
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if node := this.hashMap[key]; node != nil {
		node.Next.Val = value
		this.Get(key)
		return
	}
	node := &LinkedList{
		Key:  key,
		Val:  value,
		Next: this.listHead.Next,
	}
	this.hashMap[key] = this.listHead
	if this.listHead.Next != nil {
		this.hashMap[this.listHead.Next.Key] = node
	}
	this.listHead.Next = node
	if this.listTail == this.listHead {
		this.listTail = node
	}
	if len(this.hashMap) > this.capacity {
		key := this.listTail.Key
		this.listTail = this.hashMap[key]
		this.listTail.Next = nil
		delete(this.hashMap, key)
	}
}

func TestLRUCache(t *testing.T) {
	cache := Constructor(1)
	cache.Put(2, 1)
	fmt.Printf("cache.Get(2): %v\n", cache.Get(2))
	cache.Put(3, 2)
	fmt.Printf("cache.Get(2): %v\n", cache.Get(2))
	fmt.Printf("cache.Get(3): %v\n", cache.Get(3))
	t.Fail()
}
