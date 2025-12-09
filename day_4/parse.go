package day_4

import (
	"aoc25/parser"
	"aoc25/vec2"

	"github.com/samber/lo"
)

func parsePaperGrid(input string) []*parser.Node[bool] {
	directions := append(vec2.OrthogonalDirections, vec2.DiagonalDirections...)

	nodes := parser.ParseNodesWithDirections(input, func(item rune) bool {
		return item == '@'
	}, directions)

	return lo.Filter(nodes, func(item *parser.Node[bool], index int) bool {
		return item.Value == true
	})
}
