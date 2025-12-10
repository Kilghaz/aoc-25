package functional

import "sort"

func SortBy[T any](slice []T, iterator func(a T, b T) int) []T {
	sliceCopy := make([]T, len(slice))
	copy(sliceCopy, slice)
	sort.Slice(sliceCopy, func(i, j int) bool {
		return iterator(sliceCopy[i], sliceCopy[j]) > 0
	})
	return sliceCopy
}
