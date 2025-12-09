package math

type Numeric interface {
	~int | ~uint8 | ~uint16 | ~uint32 | ~float32 | ~float64
}
