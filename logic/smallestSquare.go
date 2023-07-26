package logic

import (
	"math"
)

func SmallestSquare(x int) int {
	elements := x * 4
	if isSquareOfNumber(elements) {
		return int(math.Pow(float64(elements), 1.0/2.0))
	}
	for i := elements; i <= 104; i++ {
		if isSquareOfNumber(i) {
			return int(math.Pow(float64(i), 1.0/2.0))
		}
	}
	return 0
}

func isSquareOfNumber(x int) bool {
	for i := 2; i <= 11; i++ {
		if i*i == x {
			return true
		}
	}
	return false
}
