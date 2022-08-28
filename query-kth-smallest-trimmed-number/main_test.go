package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

type Query struct {
	Num   string
	Index int
}

type SortQuery []Query

func (query SortQuery) Len() int {
	return len(query)
}
func (query SortQuery) Less(i, j int) bool {
	return query[i].Num < query[j].Num
}
func (query SortQuery) Swap(i, j int) {
	query[i], query[j] = query[j], query[i]
}

func Q(nums []string, query []int) int {
	sorted := make(SortQuery, 0, len(nums))
	for i, val := range nums {
		sorted = append(sorted, Query{
			Num:   val[len(val)-query[1]:],
			Index: i,
		})
	}
	sort.Stable(sorted)
	return sorted[query[0]-1].Index
}

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	ans := make([]int, 0, len(queries))
	for _, query := range queries {
		ans = append(ans, Q(nums, query))
	}
	return ans
}

func TestNum(t *testing.T) {
	assert := require.New(t)

	assert.Equal([]int{
		11, 10, 15, 1, 12, 6, 2, 15, 5, 3, 5, 12, 15, 5, 6, 7, 1, 3, 6, 14, 14, 3, 1, 8, 14, 10, 10,
	}, smallestTrimmedNumbers([]string{
		"89288488870527604910029", "36097185739782752175822", "66626740310751086142991", "39210310199276100943112", "27763774389382535382104", "38417381130871656561097", "88045540666254006395713", "95788007927435793172832", "15831923319620654311625", "45043895456202186804606", "87291364237858759125697", "88163116582250002569968", "00507332488046565482379", "57170661333341274356658", "06401271520738451116188", "21731485952024837866860",
	}, [][]int{
		{3, 17}, {10, 22}, {13, 21}, {12, 16}, {1, 23}, {10, 19}, {12, 21}, {10, 5}, {12, 9}, {12, 10}, {9, 5}, {12, 8}, {14, 6}, {5, 10}, {11, 4}, {15, 15}, {13, 9}, {1, 19}, {5, 12}, {10, 15}, {4, 18}, {4, 3}, {16, 13}, {6, 19}, {4, 18}, {2, 6}, {15, 12},
	}))

	assert.Equal([]int{
		2, 2, 1, 0,
	}, smallestTrimmedNumbers([]string{
		"102", "473", "251", "814",
	}, [][]int{
		{1, 1}, {2, 3}, {4, 2}, {1, 2},
	}))
}
