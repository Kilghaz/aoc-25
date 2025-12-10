package day_6

import "github.com/samber/lo"

func solveProblem(problem MathProblem) int {
	if problem.operator == Add {
		return lo.Sum(problem.numbers)
	}

	return lo.Reduce(problem.numbers, func(agg int, item int, _ int) int {
		return agg * item
	}, 1)
}

func Part1(input string) {
	problems := parseWorksheet(input)

	solvedProblems := lo.Map(problems, func(item MathProblem, _ int) int {
		return solveProblem(item)
	})

	println(lo.Sum(solvedProblems))
}
