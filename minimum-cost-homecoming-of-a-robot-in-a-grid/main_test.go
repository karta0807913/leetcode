package main

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) (cost int) {
	xMove := 1
	if homePos[0] < startPos[0] {
		xMove = -1
	}
	yMove := 1
	if homePos[1] < startPos[1] {
		yMove = -1
	}
	for startPos[0] != homePos[0] {
		startPos[0] += xMove
		cost += rowCosts[startPos[0]]
	}
	for startPos[1] != homePos[1] {
		startPos[1] += yMove
		cost += colCosts[startPos[1]]
	}
	return
}
