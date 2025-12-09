package slice

func ShiftN[T any](slice *[]T, n int) []T {
	local := *slice
	shifted := local[:n]
	*slice = local[n:]
	return shifted
}

func Shift[T any](slice *[]T) T {
	return ShiftN(slice, 1)[0]
}
