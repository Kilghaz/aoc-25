package day_4

import (
	"aoc25/parser"

	"github.com/samber/lo"
)

func isPaperRoll(node *parser.Node[bool], _ int) bool {
	return node.Value == true
}

func findAdjacentPaperRolls(node *parser.Node[bool]) []*parser.Node[bool] {
	return lo.Filter(node.Siblings, isPaperRoll)
}

func isPaperRollAccessible(node *parser.Node[bool], _ int) bool {
	return len(findAdjacentPaperRolls(node)) < 4
}
