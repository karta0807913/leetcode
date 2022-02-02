package main

import "testing"

func TestChannel(t *testing.T) {
	channel := make(chan int, 10)
	for i := 0; i < 10000; i++ {
		for idx := 0; idx < 10; idx++ {
			channel <- idx
		}
		for idx := 0; idx < 10; idx++ {
			idx = <-channel
		}
	}
}

func TestQueue(t *testing.T) {
	queue := make([]int, 0, 10)
	for i := 0; i < 10000; i++ {
		for idx := 0; idx < 10; idx++ {
			queue = append(queue, idx)
		}
		for idx := 0; idx < 10; idx++ {
			queue, idx = queue[1:], queue[0]
		}
	}
}
