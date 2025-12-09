package day_2

import (
	"aoc25/parser"

	"github.com/samber/lo"
)

func wholeDivisions(number int) []int {
	var divisions []int

	for i := 1; i < number; i++ {
		if number%i == 0 {
			divisions = append(divisions, i)
		}
	}

	return divisions
}

func containsRepeatedSequence(id string, _ int) bool {
	length := len(id)
	return lo.SomeBy(wholeDivisions(length), func(division int) bool {
		chunks := lo.ChunkString(id, division)
		return len(lo.Uniq(chunks)) == 1
	})
}

func Part2(input string) {
	ids := parseIds(input)

	invalidIds := lo.Filter(ids, containsRepeatedSequence)

	sum := lo.Sum(lo.Map(invalidIds, func(item string, index int) int {
		return parser.ParseInt(item)
	}))

	println(sum)
}
