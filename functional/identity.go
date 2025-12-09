package functional

func Identity[T any]() func(it T) T {
	return func(it T) T {
		return it
	}
}
