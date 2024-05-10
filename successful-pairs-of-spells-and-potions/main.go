package main

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	answer := make([]int, len(spells))
	for i, spell := range spells {
		answer[i] = len(potions) - sort.Search(len(potions), func(i int) bool {
			return success <= int64(potions[i])*int64(spell)
		})
	}
	return answer
}

/*
   4, 2, [], 4
   [0,0,1,0]
*/
