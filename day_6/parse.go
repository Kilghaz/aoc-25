package day_6

import (
	"aoc25/parser"
	"aoc25/slice"
	"regexp"
	"strings"

	"github.com/samber/lo"
)

func parseWorksheet(input string) []MathProblem {
	spaces := regexp.MustCompile("\\s+")
	lines := strings.Split(input, "\n")
	numbers := lo.Map(lines[:len(lines)-1], func(line string, index int) []int {
		return lo.Map(spaces.Split(strings.TrimSpace(line), -1), func(numerStr string, index int) int {
			return parser.ParseInt(numerStr)
		})
	})
	operators := spaces.Split(slice.Last(lines), -1)

	return lo.Map(operators, func(operator string, index int) MathProblem {
		problemNumbers := lo.Map(numbers, func(item []int, _ int) int {
			return item[index]
		})
		return MathProblem{numbers: problemNumbers, operator: operator}
	})
}
