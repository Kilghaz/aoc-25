package slice

import "github.com/samber/lo"

func PopN[T any](slice *[]T, n int) []T {
	local := *slice
	popped := local[len(local)-n:]
	*slice = local[:len(local)-n]
	return lo.Reverse(popped)
}

func Pop[T any](slice *[]T) []T {
	return PopN(slice, 1)
}
