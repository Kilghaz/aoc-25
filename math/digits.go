package math

func CountDigits(num int) int {
	return CountDigitsBase(num, 10)
}

func CountDigitsBase(num int, base int) int {
	if num == 0 {
		return 1
	}
	digits := 0
	for num != 0 {
		num /= base
		digits++
	}
	return digits
}
