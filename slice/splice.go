package slice

func Splice[T any](slice *[]T, start, length int) []T {
	local := *slice
	spliced := local[start : start+length]
	result := make([]T, 0)
	result = append(result, local[:start]...)
	result = append(result, local[start+length:]...)
	*slice = result
	return spliced
}
