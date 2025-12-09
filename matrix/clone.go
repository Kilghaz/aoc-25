package matrix

func Clone[T any](matrix [][]T) [][]T {
	rows := len(matrix)
	cols := len(matrix[0])
	cloned := make([][]T, rows)
	data := make([]T, rows*cols)
	for i := range matrix {
		start := i * cols
		end := start + cols
		cloned[i] = data[start:end:end]
		copy(cloned[i], matrix[i])
	}
	return cloned
}
