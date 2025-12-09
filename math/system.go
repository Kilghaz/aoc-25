package math

func ToBase(number, base int) []int {
	if number == 0 {
		return []int{0}
	}

	digits := make([]int, 0)
	for number > 0 {
		remainder := number % base
		digits = append(digits, remainder)
		number /= base
	}
	return digits
}
