package math

import "math"

func Difference(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}
