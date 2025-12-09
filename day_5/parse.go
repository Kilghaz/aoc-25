package day_5

import (
	"aoc25/parser"
	"strings"

	"github.com/samber/lo"
)

func parseFreshnessRange(rangeInput string) FreshnessRange {
	parts := strings.Split(rangeInput, "-")
	start := parser.ParseInt(parts[0])
	end := parser.ParseInt(parts[1])
	return FreshnessRange{
		start: start,
		end:   end,
	}
}

func parseIngredientDatabase(input string) IngredientDatabase {
	parts := strings.Split(input, "\n\n")
	available := lo.Map(strings.Split(parts[1], "\n"), func(item string, index int) int {
		return parser.ParseInt(item)
	})
	freshnessRanges := lo.Map(strings.Split(parts[0], "\n"), func(line string, index int) FreshnessRange {
		return parseFreshnessRange(line)
	})

	return IngredientDatabase{
		freshnessRanges: freshnessRanges,
		available:       available,
	}
}
