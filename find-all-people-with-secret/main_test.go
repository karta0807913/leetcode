package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

type UnionFind map[int]int

func (u UnionFind) FindParent(children int) int {
	parent, ok := u[children]
	if !ok {
		u[children] = children
		return children
	}
	if parent != children {
		parent = u.FindParent(parent)
		u[children] = parent
	}
	return parent
}

func (u UnionFind) JoinGroup(parent, children int, knowsPpl map[int]struct{}) {
	parent, children = u.FindParent(parent), u.FindParent(children)
	if _, ok := knowsPpl[children]; ok {
		parent, children = children, parent
	}
	u[children] = parent
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	knowsPpl := make(map[int]struct{})
	knowsPpl[firstPerson] = struct{}{}
	knowsPpl[0] = struct{}{}

	for index := 0; index < len(meetings); index++ {
		meeting := meetings[index]
		time := meeting[2]
		uFind := make(UnionFind)
		uFind.JoinGroup(meeting[0], meeting[1], knowsPpl)
		for index+1 != len(meetings) && meetings[index+1][2] == time {
			meeting = meetings[index+1]
			uFind.JoinGroup(meeting[0], meeting[1], knowsPpl)
			index++
		}
		for person := range uFind {
			if _, ok := knowsPpl[uFind.FindParent(person)]; ok {
				knowsPpl[person] = struct{}{}
			}
		}
	}
	ans := make([]int, 0, len(knowsPpl)+1)
	for person := range knowsPpl {
		ans = append(ans, person)
	}
	return ans
}

func TestFindAllPeople(t *testing.T) {
	assert := require.New(t)
	assert.ElementsMatch(
		[]int{0, 1, 2, 3, 5},
		findAllPeople(6, [][]int{
			{1, 2, 5}, {2, 3, 8}, {1, 5, 10},
		}, 1),
	)
}
