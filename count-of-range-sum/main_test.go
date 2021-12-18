package countofrangesum

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var depth = 0

func sumArr(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func countRangeSum(nums []int, lower int, upper int) int {
	depth += 1
	if depth > 20 {
		os.Exit(127)
	}
	fmt.Println(nums)
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if lower <= nums[0] && nums[0] <= upper {
			return 1
		}
		return 0
	}
	if len(nums) == 2 {
		sum := nums[0] + nums[1]
		count := 0
		if lower <= nums[0] && nums[0] <= upper {
			count += 1
		}
		if lower <= nums[1] && nums[1] <= upper {
			count += 1
		}
		if lower <= sum && sum <= upper {
			count += 1
		}
		return count
	}
	np := make([]int, 0)
	sum := 0
	bp := 0
	count := 0
	for idx := range nums {
		sum += nums[idx]
		if sum < lower || upper < sum {
			count += countRangeSum(nums[bp:idx+1], lower, upper)
			np = append(np, sum)
			bp = idx
		}
	}
	if bp != 0 {
		count += countRangeSum(nums[bp:], lower, upper)
	}
	if lower <= sum && sum <= upper {
		count += 1
	}
	np = append(np, sum)
	fmt.Println("A", np)
	if len(np) != len(nums) {
		count += countRangeSum(np, lower, upper)
	}
	return count
}

func TestResult(t *testing.T) {
	// assert.Equal(t, 0, countRangeSum([]int{5, -11}, -2, 2))
	// assert.Equal(t, 1, countRangeSum([]int{5, -1}, -2, 2))
	// assert.Equal(t, 3, countRangeSum([]int{-2, 5, -1}, -2, 2))
	// assert.Equal(t, 4, countRangeSum([]int{-2, 5, -3, 3}, -2, 2))
	// assert.Equal(t, 3, countRangeSum([]int{-2, 5, -3}, -2, 2))
	// assert.Equal(t, 3, countRangeSum([]int{3, -3, 0}, -2, 2))
	assert.Equal(t, 6, countRangeSum([]int{3, -5, 3, 0}, -2, 2))
	// assert.Equal(t, 3, countRangeSum([]int{-5, 3, 0}, -2, 2))
	// assert.Equal(t, 3, countRangeSum([]int{3, -5, -3, 0}, -2, 2))
}
