package slice

func Pad[T any](slice []T, length int, pad T) []T {
	padded := make([]T, length)
	for i := range padded {
		padded[i] = pad
	}
	for i := range slice {
		padded[i] = slice[i]
	}
	return padded
}
