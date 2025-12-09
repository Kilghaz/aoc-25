package parser

import (
	"aoc25/matrix"
	"github.com/samber/lo"
	"strings"
)

func ParseGrid(input string) [][]rune {
	return matrix.Invert(lo.Map(strings.Split(input, "\n"), func(row string, _ int) []rune {
		return []rune(row)
	}))
}
