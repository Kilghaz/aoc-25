package matrix

func Dimension[T any](matrix [][]T) (int, int) {
	return len(matrix), len(matrix[0])
}
