package main

import "math"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	T := minutesToTest/minutesToDie + 1
	pig := 1
	for math.Pow(float64(T), float64(pig)) < float64(buckets) {
		pig += 1
	}
	return pig
}
