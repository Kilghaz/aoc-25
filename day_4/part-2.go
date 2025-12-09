package day_4

import (
	"aoc25/parser"

	"github.com/samber/lo"
)

func Part2(input string) {
	nodes := parsePaperGrid(input)

	removedPaperRolls := 0

	for {
		accessiblePaperRolls := lo.Filter(nodes, isPaperRollAccessible)

		if len(accessiblePaperRolls) == 0 {
			break
		}

		removedPaperRolls += len(accessiblePaperRolls)
		lo.ForEach(accessiblePaperRolls, func(item *parser.Node[bool], index int) {
			item.Value = false // remove paper roll
		})

		nodes = lo.Filter(nodes, isPaperRoll)
	}

	println(removedPaperRolls)
}
