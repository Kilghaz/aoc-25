package day_5

type FreshnessRange struct {
	start int
	end   int
}

type IngredientDatabase struct {
	freshnessRanges []FreshnessRange
	available       []int
}
