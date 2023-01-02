package main

import (
	"container/heap"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

type ScheduleHeap []Schedule

func (h ScheduleHeap) Len() int           { return len(h) }
func (h ScheduleHeap) Less(i, j int) bool { return h[i].EndTime < h[j].EndTime }
func (h ScheduleHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ScheduleHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Schedule))
}

func (h *ScheduleHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Job struct {
	StartTime int
	EndTime   int
	Profit    int
}

func (h *ScheduleHeap) SquashValues(time, maxVal int) int {
	if len(*h) == 0 {
		return maxVal
	}
	for tmp := (*h)[0]; tmp.EndTime <= time; tmp = (*h)[0] {
		heap.Pop(h)
		if tmp.Value > maxVal {
			maxVal = tmp.Value
		}
		if len(*h) == 0 {
			return maxVal
		}
	}
	return maxVal
}

type Schedule struct {
	EndTime int
	Value   int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := make([]Job, 0, len(startTime))
	for i := range startTime {
		jobs = append(jobs, Job{
			StartTime: startTime[i],
			EndTime:   endTime[i],
			Profit:    profit[i],
		})
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].StartTime < jobs[j].StartTime
	})
	schedule := make(ScheduleHeap, 0)
	maxValue := 0
	for _, job := range jobs {
		// fmt.Printf("job: %v\n", job)
		// fmt.Printf("schedule: %v\n", schedule)
		maxValue = schedule.SquashValues(job.StartTime, maxValue)
		heap.Push(&schedule, Schedule{
			EndTime: job.EndTime,
			Value:   maxValue + job.Profit,
		})
		// fmt.Printf("schedule: %v\n", schedule)
	}
	maxValue = schedule.SquashValues(1e9+1, maxValue)
	return maxValue
}

func TestScheduling(t *testing.T) {
	assert := require.New(t)
	assert.Equal(120, jobScheduling(
		[]int{1, 2, 3, 3},
		[]int{3, 4, 5, 6},
		[]int{50, 10, 40, 70},
	))
}
