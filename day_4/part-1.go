package day_4

import (
	"github.com/samber/lo"
)

func Part1(input string) {
	nodes := parsePaperGrid(input)
	accessiblePaperRolls := lo.Filter(nodes, isPaperRollAccessible)
	println(len(accessiblePaperRolls))
}
