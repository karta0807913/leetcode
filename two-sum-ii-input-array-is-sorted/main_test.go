package main

func binarySearch(numbers []int, target int) int {
	left, right := 0, len(numbers)-1
	for left <= right {
		mid := (left + right) / 2
		if numbers[mid] == target {
			return mid
		} else if numbers[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// func twoSum(numbers []int, target int) []int {
// 	for i, num := range numbers {
// 		idx := binarySearch(numbers[i+1:], target-num)
// 		if idx != -1 {
// 			fmt.Printf("numbers[idx]: %v\n", numbers[idx+i+1])
// 			return []int{i + 1, idx + 2 + i}
// 		}
// 	}
// 	return []int{0, 0}
// }

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left <= right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
