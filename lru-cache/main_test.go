package main

import (
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/require"
)

type LinkList struct {
	Key       int
	Val       int
	Destroyed bool
	Next      *LinkList
}

type LRUCache struct {
	cache                    map[int]*LinkList
	head                     *LinkList
	tail                     **LinkList
	listSize, size, Capacity int
}

func Constructor(capacity int) LRUCache {
	head := &LinkList{}
	return LRUCache{
		cache:    make(map[int]*LinkList),
		head:     head,
		tail:     &head.Next,
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	node.Destroyed = true
	this.listSize += 1
	this.cache[key] = this.appendList(key, node.Val)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		node.Destroyed = true
	} else {
		this.size += 1
	}
	this.cache[key] = this.appendList(key, value)
	if this.size > this.Capacity {
		this.pop()
	}
}

func (this *LRUCache) appendList(key, val int) *LinkList {
	this.listSize += 1
	node := &LinkList{
		Val: val,
		Key: key,
	}
	*this.tail = node
	this.tail = &node.Next
	if this.listSize > this.Capacity*this.Capacity {
		this.cleanList()
	}
	return node
}

func (this *LRUCache) cleanList() {
	listSize := 0
	pointer := this.head
	for pointer.Next != nil {
		if pointer.Next.Destroyed {
			pointer.Next = pointer.Next.Next
		} else {
			listSize += 1
			pointer = pointer.Next
		}
		this.tail = &pointer.Next
	}
	this.listSize = listSize
}

func (this *LRUCache) pop() {
	if this.head.Next != nil {
		for this.head.Next.Destroyed {
			this.head = this.head.Next
			this.listSize -= 1
		}
		if this.head.Next != nil {
			delete(this.cache, this.head.Next.Key)
			this.head = this.head.Next
		}
		this.size -= 1
	}
}

func TestLRUCahce(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)
	cache.Put(3, 3)
	assert.Equal(t, -1, cache.Get(2))

	cache = Constructor(1)
	assert.Equal(t, -1, cache.Get(6))
	assert.Equal(t, -1, cache.Get(8))
	cache.Put(12, 1)
	assert.Equal(t, -1, cache.Get(2))
	fmt.Println(cache.cache)
	cache.Put(15, 11)
	fmt.Println(cache.cache)
	cache.Put(5, 2)
	fmt.Println(cache.cache)
	cache.Put(1, 15)
	fmt.Println(cache.cache, cache.size, cache.listSize)
	cache.Put(4, 2)
	fmt.Println(cache.cache)
	assert.Equal(t, -1, cache.Get(5))
	cache.Put(15, 15)
}
