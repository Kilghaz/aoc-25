package day_2

import (
	"aoc25/parser"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func expandRange(rangeInput string) []string {
	parts := strings.Split(rangeInput, "-")
	start := parser.ParseInt(parts[0])
	end := parser.ParseInt(parts[1])
	return lo.Map(lo.RangeFrom(start, end-start+1), func(item int, _index int) string {
		return strconv.Itoa(item)
	})
}

func parseIds(input string) []string {
	return lo.FlatMap(strings.Split(input, ","), func(item string, _index int) []string {
		return expandRange(item)
	})
}
