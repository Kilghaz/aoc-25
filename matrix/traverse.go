package matrix

type TraversalOffset struct {
	startX int
	startY int
	endX   int
	endY   int
}

func SymmetricOffset(n int) TraversalOffset {
	return TraversalOffset{
		startX: n,
		startY: n,
		endX:   n,
		endY:   n,
	}
}

func TraverseOffset[T any](matrix [][]T, offset TraversalOffset, cb func(item T, x int, y int)) {
	for x := offset.startX; x < len(matrix)-offset.endX; x++ {
		for y := offset.startY; y < len(matrix[x])-offset.endY; y++ {
			cb(matrix[x][y], x, y)
		}
	}
}

func TraverseSymmetricallyOffset[T any](matrix [][]T, offset int, cb func(item T, x int, y int)) {
	TraverseOffset(matrix, SymmetricOffset(offset), cb)
}

func Traverse[T any](matrix [][]T, cb func(item T, x int, y int)) {
	TraverseOffset(matrix, SymmetricOffset(0), cb)
}
