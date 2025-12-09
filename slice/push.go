package slice

func Push[T any](slice *[]T, value T) *[]T {
	*slice = append(*slice, value)
	return slice
}
