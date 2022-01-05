package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type UnionSet struct {
	islandMap map[byte]byte
}

func (set *UnionSet) Find(x byte) byte {
	parent, ok := set.islandMap[x]
	if !ok {
		set.islandMap[x] = x
		return x
	}
	if parent != x {
		set.islandMap[x] = set.Find(parent)
		return set.islandMap[x]
	}
	return x
}

func (set *UnionSet) Union(x, y byte) {
	x, y = set.Find(x), set.Find(y)
	if x != y {
		set.islandMap[x] = y
	}
}

func equationsPossible(equations []string) bool {
	unionSet := UnionSet{
		islandMap: make(map[byte]byte),
	}
	neqList := make([]byte, 0)
	for _, eq := range equations {
		switch eq[1] {
		case '!':
			if eq[0] == eq[3] {
				return false
			}
			neqList = append(neqList, eq[0], eq[3])
		case '=':
			unionSet.Union(eq[0], eq[3])
		}
	}
	for idx := 0; idx < len(neqList); idx += 2 {
		left := unionSet.Find(neqList[idx])
		right := unionSet.Find(neqList[idx+1])
		if left == right {
			return false
		}
	}
	return true
}

func TestEquations(t *testing.T) {
	assert.False(t, equationsPossible([]string{"a!=a"}))
	assert.True(t, equationsPossible([]string{"a==a"}))
	assert.False(t, equationsPossible([]string{"a==b", "b!=a"}))
	assert.True(t, equationsPossible([]string{"a==b", "b==a"}))
	assert.True(t, equationsPossible([]string{"e==e", "d!=e", "c==d", "d!=e"}))
}
