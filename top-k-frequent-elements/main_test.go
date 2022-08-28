package main

func topKFrequent(nums []int, k int) []int {
	bucket := make([][]int, len(nums)+1)
	appearedTimes := make(map[int]int)

	for _, num := range nums {
		appearedTimes[num] += 1
	}
	for num, times := range appearedTimes {
		bucket[times] = append(bucket[times], num)
	}
	result := make([]int, 0, k)
	for i := len(bucket) - 1; i >= 0; i-- {
		result = append(result, bucket[i]...)
		if len(result) >= k {
			return result[:k]
		}
	}
	return result[:k]
}
