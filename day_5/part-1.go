package day_5

import "github.com/samber/lo"

func isInRange(freshnessRange FreshnessRange, ingredient int) bool {
	return ingredient >= freshnessRange.start && ingredient <= freshnessRange.end
}

func isFresh(db IngredientDatabase, ingredient int) bool {
	return lo.SomeBy(db.freshnessRanges, func(item FreshnessRange) bool {
		return isInRange(item, ingredient)
	})
}

func Part1(input string) {
	db := parseIngredientDatabase(input)

	freshAndAvailableIngredients := lo.Filter(db.available, func(item int, index int) bool {
		return isFresh(db, item)
	})

	println(len(freshAndAvailableIngredients))
}
