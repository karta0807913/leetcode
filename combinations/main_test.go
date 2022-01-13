package main

import (
	"fmt"
	"testing"
)

func combine(n int, k int) (answer [][]int) {
	current := make([]int, k)
	for i := 0; i < k; i++ {
		current[i] = n - i
	}
	answer = append(answer, append([]int{}, current...))
	for c := 0; c < 10; c++ {
		carry := -1
		idx := k - 1
		flag := false
		for ; idx >= 0 && carry != 0; idx-- {
			current[idx] += carry
			carry = 0
			// fmt.Println(current, idx, k-1-idx)
			if current[idx] == (k-1)-idx {
				flag = true
				carry = -1
			}
		}
		if carry != 0 {
			break
		}
		idx += 1
		// fmt.Printf("idx: %v, current: %v\n", idx, current)
		for ; idx < k-1 && flag; idx++ {
			current[idx+1] = current[idx] - 1
			if current[idx+1] == 0 {
				return
			}
		}

		// fmt.Printf("carry: %v\n", carry)
		answer = append(answer, append([]int{}, current...))
	}
	return
}

func TestCombine(t *testing.T) {
	// fmt.Println(combine(4, 2))
	fmt.Println(combine(4, 3))
	// fmt.Println(combine(4, 4))
	t.Fail()
}
