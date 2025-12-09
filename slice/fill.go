package slice

import "slices"

func Fill[T any](slice []T, value T) []T {
	cloned := slices.Clone(slice)
	for i := 0; i < len(cloned); i++ {
		cloned[i] = value
	}
	return cloned
}
