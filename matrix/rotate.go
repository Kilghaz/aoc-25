package matrix

import "slices"

func Rotate[T any](matrix [][]T) [][]T {
	rows := len(matrix)
	cols := len(matrix[0])

	rotated := make([][]T, cols)
	for i := range rotated {
		rotated[i] = make([]T, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-1-i] = matrix[i][j]
		}
	}

	return rotated
}

func RotateN[T any](matrix [][]T, n int) [][]T {
	rotated := slices.Clone(matrix)
	for i := 0; i < n; i++ {
		rotated = Rotate(rotated)
	}
	return rotated
}
