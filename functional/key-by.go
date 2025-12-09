package functional

import "github.com/samber/lo"

func KeyBy[T any, K comparable](input []T, iterator func(item T, index int) K) map[K][]T {
	result := make(map[K][]T)
	for index, item := range input {
		key := iterator(item, index)
		if !lo.HasKey(result, key) {
			result[key] = []T{item}
			continue
		}
		result[key] = append(result[key], item)
	}
	return result
}
