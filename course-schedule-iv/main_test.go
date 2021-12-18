package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func buildDatabase(target int, database [][]bool, searched []bool) {
	if searched[target] {
		return
	}
	for dstIdx := range database[target] {
		if database[target][dstIdx] {
			buildDatabase(dstIdx, database, searched)
			for idx := range database[dstIdx] {
				database[target][idx] = database[target][idx] || database[dstIdx][idx]
			}
		}
	}
	searched[target] = true
}

func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
	queryDatabase := make([][]bool, n)
	for idx := range queryDatabase {
		queryDatabase[idx] = make([]bool, n)
	}
	for _, val := range prerequisites {
		//           target   pretarget
		queryDatabase[val[1]][val[0]] = true
	}
	searched := make([]bool, n)
	for target := range queryDatabase {
		if searched[target] {
			continue
		}
		buildDatabase(target, queryDatabase, searched)
	}
	answer := make([]bool, len(queries))
	for idx, query := range queries {
		answer[idx] = queryDatabase[query[1]][query[0]]
	}
	return answer
}

func TestBuildDatabase(t *testing.T) {
	n := 3
	database := make([][]bool, n)
	for idx := range database {
		database[idx] = make([]bool, n)
	}
	database[0][1] = true
	database[1][2] = true
	searched := make([]bool, n)
	for target := range database {
		buildDatabase(target, database, searched)
	}
	fmt.Println(database)
	assert.Equal(t, [][]bool{
		{false, true, true},
		{false, false, true},
		{false, false, false},
	}, database, nil)

	assert.Equal(t, []bool{false, true}, checkIfPrerequisite(2, [][]int{{1, 0}}, [][]int{{0, 1}, {1, 0}}), nil)
	assert.Equal(t, []bool{false, false}, checkIfPrerequisite(2, [][]int{}, [][]int{{1, 0}, {0, 1}}), nil)
	assert.Equal(t, []bool{false, true}, checkIfPrerequisite(5, [][]int{{1, 0}, {2, 0}}, [][]int{{0, 1}, {2, 0}}), nil)
	assert.Equal(t, []bool{true, false, true, false}, checkIfPrerequisite(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, [][]int{{0, 4}, {4, 0}, {1, 3}, {3, 0}}), nil)
}
