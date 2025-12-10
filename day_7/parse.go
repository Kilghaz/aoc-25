package day_7

import (
	"aoc25/functional"
	"aoc25/parser"
	"aoc25/vec2"

	"github.com/samber/lo"
)

type TachyonManifold struct {
	start    vec2.Vector2[int]
	splitter []vec2.Vector2[int]
}

func parseTachyonManifold(input string) TachyonManifold {
	nodes := parser.ParseNodes(input, func(item rune) rune {
		return item
	})

	startNode, _ := lo.Find(nodes, func(node *parser.Node[rune]) bool {
		return node.Value == 'S'
	})

	splitterNodes := lo.Filter(nodes, func(node *parser.Node[rune], _ int) bool {
		return node.Value == '^'
	})

	splitter := lo.Map(splitterNodes, func(item *parser.Node[rune], _ int) vec2.Vector2[int] {
		return item.Position
	})

	return TachyonManifold{
		start: startNode.Position,
		splitter: functional.SortBy(splitter, func(a vec2.Vector2[int], b vec2.Vector2[int]) int {
			return b.Y - a.Y
		}),
	}
}
