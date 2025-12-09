package math

import "math"

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func Log10(num int) int {
	return int(math.Log10(float64(num)))
}

func Abs(num int) int {
	return int(math.Abs(float64(num)))
}
