package matrix

func Invert[T any](matrix [][]T) [][]T {
	inverted := make([][]T, len(matrix))
	for j := 0; j < len(matrix[0]); j++ {
		for i := 0; i < len(matrix); i++ {
			inverted[j] = append(inverted[j], matrix[i][j])
		}
	}
	return inverted
}
