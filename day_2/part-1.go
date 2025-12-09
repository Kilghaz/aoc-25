package day_2

import (
	"aoc25/parser"

	"github.com/samber/lo"
)

func containsSingleRepeatedSequence(id string, _ int) bool {
	if len(id)%2 != 0 {
		return false
	}

	return lo.SomeBy([]int{2}, func(division int) bool {
		chunks := lo.ChunkString(id, len(id)/division)
		return len(lo.Uniq(chunks)) == 1
	})
}

func Part1(input string) {
	ids := parseIds(input)

	invalidIds := lo.Filter(ids, containsSingleRepeatedSequence)

	sum := lo.Sum(lo.Map(invalidIds, func(item string, index int) int {
		return parser.ParseInt(item)
	}))

	println(sum)
}
