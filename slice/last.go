package slice

func Last[T any](slice []T) T {
	return slice[len(slice)-1]
}
