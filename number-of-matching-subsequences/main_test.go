package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func numMatchingSubseq(s string, words []string) (num int) {
	indexSet := make([][]int, 26)
	for i, b := range s {
		c := &indexSet[b-'a']
		*c = append(*c, i)
	}
	for _, word := range words {
		leftBound := -1
		found := true
		for _, char := range word {
			c := indexSet[char-'a']
			left := 0
			right := len(c)
			// check if index inside [leftBound, len(s))
			for left < right {
				mid := (left + right) / 2
				if leftBound < c[mid] {
					right = mid
				} else {
					left = mid + 1
				}
			}
			if right != len(c) {
				leftBound = c[right]
			} else {
				found = false
				break
			}
		}
		if found {
			num += 1
		}
	}
	return
}

func TestSubseq(t *testing.T) {
	assert.Equal(t, 2, numMatchingSubseq("abcbcac", []string{
		"b", "abc",
	}))
	assert.Equal(t, 5, numMatchingSubseq("rwpddkvbnnuglnagtvamxkqtwhqgwbqgfbvgkwyuqkdwhzudsxvjubjgloeofnpjqlkdsqvruvabjrikfwronbrdyyjnakstqjac", []string{
		"lnagtva",
		"kvbnnuglnagtvamxkqtwhqgwbqgfbvgkwyuqkdwhzudsxvju",
		"rwpddkvbnnugln",
		"gloeofnpjqlkdsqvruvabjrikfwronbrdyyj",
		"vbgeinupkvgmgxeaaiuiyojmoqkahwvbpwugdainxciedbdkos",
		"mspuhbykmmumtveoighlcgpcapzczomshiblnvhjzqjlfkpina",
		"rgmliajkiknongrofpugfgajedxicdhxinzjakwnifvxwlokip",
		"fhepktaipapyrbylskxddypwmuuxyoivcewzrdwwlrlhqwzikq",
		"qatithxifaaiwyszlkgoljzkkweqkjjzvymedvclfxwcezqebx",
		"wpddkvbnn",
	}))
	assert.Equal(t, 1, numMatchingSubseq("abdcdbez", []string{
		"abdz",
	}))
}
