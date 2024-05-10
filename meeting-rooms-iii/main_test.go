package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type TimeHeap []TimeOperation

func (h TimeHeap) Len() int { return len(h) }
func (h TimeHeap) Less(i, j int) bool {
	if h[i].OperationTime() == h[j].OperationTime() {
		return h[i].Type() < h[j].Type()
	}
	return h[i].OperationTime() < h[j].OperationTime()
}
func (h TimeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *TimeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(TimeOperation))
}

func (h *TimeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type TimeOperation interface {
	OperationTime() int
	Type() int
	Do()
}

type RoomEndTime struct {
	EndAt     int
	RoomIndex int
	roomHeap  heap.Interface
}

func (t *RoomEndTime) OperationTime() int {
	return t.EndAt
}

func (t *RoomEndTime) Type() int {
	return 0
}

func (t *RoomEndTime) Do() {
	heap.Push(t.roomHeap, t.RoomIndex)
}

type MeetingTime struct {
	StartTime   int
	EndTime     int
	meetingHeap heap.Interface
}

func (t *MeetingTime) OperationTime() int {
	return t.StartTime
}

func (t *MeetingTime) Type() int {
	return 1
}

func (t *MeetingTime) Do() {
	heap.Push(t.meetingHeap, t)
}

func mostBooked(n int, meetings [][]int) int {
	record := make([]int, n)
	var roomHeap IntHeap
	for i := 0; i < n; i++ {
		roomHeap = append(roomHeap, i)
	}
	heap.Init(&roomHeap)

	var delayedMeeting TimeHeap
	var operations TimeHeap
	for _, meeting := range meetings {
		operations = append(operations, &MeetingTime{
			StartTime:   meeting[0],
			EndTime:     meeting[1],
			meetingHeap: &delayedMeeting,
		})
	}
	heap.Init(&operations)

	for len(operations) != 0 {
		operation := heap.Pop(&operations).(TimeOperation)
		operation.Do()
		fmt.Printf("operation: %+v\n", operation)
		if len(operations) != 0 && operations[0].OperationTime() == operation.OperationTime() {
			continue
		}
		for !(len(roomHeap) == 0 || len(delayedMeeting) == 0) {
			room := heap.Pop(&roomHeap).(int)
			meeting := heap.Pop(&delayedMeeting).(*MeetingTime)
			heap.Push(&operations, &RoomEndTime{
				EndAt:     operation.OperationTime() + meeting.EndTime - meeting.StartTime,
				RoomIndex: room,
				roomHeap:  &roomHeap,
			})
			record[room]++
			fmt.Println(room, record)
		}
	}
	maxVal := 0
	maxRoom := 0
	for i, v := range record {
		if v > maxVal {
			maxVal = v
			maxRoom = i
		}
	}
	fmt.Println(record)
	return maxRoom
}
