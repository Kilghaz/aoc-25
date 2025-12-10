package day_6

import "github.com/samber/lo"

type MathProblem struct {
	numbers  []int
	operator rune
}

const (
	Add      = '+'
	Multiply = '*'
)

func solveProblem(problem MathProblem) int {
	if problem.operator == Add {
		return lo.Sum(problem.numbers)
	}

	return lo.Reduce(problem.numbers, func(agg int, item int, _ int) int {
		return agg * item
	}, 1)
}
