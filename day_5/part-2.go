package day_5

import (
	"aoc25/functional"
	"aoc25/slice"
	"slices"

	"github.com/samber/lo"
)

func doRangesOverlap(a FreshnessRange, b FreshnessRange) bool {
	return a.start >= b.start && a.start <= b.end ||
		a.end >= b.start && a.end <= b.end ||
		b.start >= a.start && b.start <= a.end ||
		b.end >= a.start && b.end <= a.end
}

func mergeRanges(ranges []FreshnessRange, freshnessRange FreshnessRange) []FreshnessRange {
	if len(ranges) == 0 {
		return []FreshnessRange{freshnessRange}
	}

	ranges = functional.SortBy(ranges, func(a FreshnessRange, b FreshnessRange) int {
		return a.start - b.start
	})

	var overlappingIndices []int
	overlappingRanges := lo.Filter(ranges, func(item FreshnessRange, index int) bool {
		doOverlap := doRangesOverlap(item, freshnessRange)
		if doOverlap {
			overlappingIndices = append(overlappingIndices, index)
		}
		return doOverlap
	})

	if len(overlappingRanges) == 0 {
		return append(ranges, freshnessRange)
	}

	ranges = slices.Delete(ranges, overlappingIndices[0], slice.Last(overlappingIndices)+1)

	overlappingRanges = append(overlappingRanges, freshnessRange)
	start := lo.Min(lo.Map(overlappingRanges, func(item FreshnessRange, index int) int {
		return item.start
	}))
	end := lo.Max(lo.Map(overlappingRanges, func(item FreshnessRange, index int) int {
		return item.end
	}))

	return append(ranges, FreshnessRange{start: start, end: end})
}

func countIngredients(freshnessRange FreshnessRange) int {
	return freshnessRange.end - freshnessRange.start + 1
}

func Part2(input string) {
	db := parseIngredientDatabase(input)

	var globalFreshnessRanges []FreshnessRange

	lo.ForEach(db.freshnessRanges, func(item FreshnessRange, index int) {
		globalFreshnessRanges = mergeRanges(globalFreshnessRanges, item)
	})

	ingredientsCount := lo.Sum(lo.Map(globalFreshnessRanges, func(item FreshnessRange, _ int) int {
		return countIngredients(item)
	}))

	println(ingredientsCount)
}
