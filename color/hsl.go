package color

import "math"

func ToHSL(r, g, b uint8) (float64, float64, float64) {
	var h, s, l float64

	rf := float64(r) / 255
	gf := float64(g) / 255
	bf := float64(b) / 255

	maxValue := math.Max(math.Max(rf, gf), bf)
	minValue := math.Min(math.Min(rf, gf), bf)

	l = (maxValue + minValue) / 2

	delta := maxValue - minValue
	if delta == 0 {
		return 0, 0, l
	}

	if l < 0.5 {
		s = delta / (maxValue + minValue)
	} else {
		s = delta / (2 - maxValue - minValue)
	}

	r2 := (((maxValue - rf) / 6) + (delta / 2)) / delta
	g2 := (((maxValue - gf) / 6) + (delta / 2)) / delta
	b2 := (((maxValue - bf) / 6) + (delta / 2)) / delta
	switch {
	case rf == maxValue:
		h = b2 - g2
	case gf == maxValue:
		h = (1.0 / 3.0) + r2 - b2
	case bf == maxValue:
		h = (2.0 / 3.0) + g2 - r2
	}

	switch {
	case h < 0:
		h += 1
	case h > 1:
		h -= 1
	}

	return h, s, l
}
