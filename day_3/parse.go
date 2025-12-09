package day_3

import (
	"aoc25/parser"
	"strings"

	"github.com/samber/lo"
)

func parseBatteriesBanks(input string) [][]int {
	return lo.Map(strings.Split(input, "\n"), func(bank string, index int) []int {
		return lo.Map(lo.ChunkString(bank, 1), func(battery string, index int) int {
			return parser.ParseInt(battery)
		})
	})
}
