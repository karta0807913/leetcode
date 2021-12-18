package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Job struct {
	index int
	task  []int
}

func getOrder(tasks [][]int) []int {
	if len(tasks) == 0 {
		return []int{}
	}
	job := make([]Job, 0, len(tasks))
	for idx, val := range tasks {
		job = append(job, Job{
			index: idx,
			task:  val,
		})
	}
	sort.Slice(job, func(i, j int) bool {
		if job[i].task[0] == job[j].task[0] {
			return job[i].task[1] < job[j].task[1]
		}
		return job[i].task[0] < job[j].task[0]
	})
	available := []Job{}
	result := []int{}
	currentTime := 0
	for len(job) != 0 {
		if len(available) == 0 {
			available = append(available, job[0])
			currentTime = job[0].task[0]
			job = job[1:]
		}
		result = append(result, available[0].index)
		currentTime += available[0].task[1]
		available = available[1:]
		idx := 0
		for idx = 0; idx < len(job); idx++ {
			if job[idx].task[0] <= currentTime {
				available = append(available, job[idx])
			} else {
				break
			}
		}
		job = job[idx:]
		sort.Slice(available, func(i, j int) bool {
			if available[i].task[1] == available[j].task[1] {
				return available[i].index < available[j].index
			}
			return available[i].task[1] < available[j].task[1]
		})
	}
	for _, val := range available {
		result = append(result, val.index)
	}
	return result
}

func TestCPU(t *testing.T) {
	// assert.Equal(t, []int{0, 2, 3, 1}, getOrder([][]int{
	// 	{1, 2},
	// 	{2, 4},
	// 	{3, 2},
	// 	{4, 1},
	// }), nil)

	// assert.Equal(t, []int{0, 2, 3, 1}, getOrder([][]int{
	// 	{1, 2},
	// 	{2, 4},
	// 	{3, 2},
	// 	{3, 2},
	// }), nil)

	// assert.Equal(t, []int{4, 3, 2, 0, 1}, getOrder([][]int{
	// 	{7, 10},
	// 	{7, 12},
	// 	{7, 5},
	// 	{7, 4},
	// 	{7, 2},
	// }), nil)

	assert.Equal(t, []int{6, 1, 2, 9, 4, 10, 0, 11, 5, 13, 3, 8, 12, 7}, getOrder([][]int{
		{19, 13},
		{16, 9},
		{21, 10},
		{32, 25},
		{37, 4},
		{49, 24},
		{2, 15},
		{38, 41},
		{37, 34},
		{33, 6},
		{45, 4},
		{18, 18},
		{46, 39},
		{12, 24},
	}), nil)
}
