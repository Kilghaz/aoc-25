package day_6

import "github.com/samber/lo"

func Part1(input string) {
	problems := parseWorksheet(input)

	solvedProblems := lo.Map(problems, func(item MathProblem, _ int) int {
		return solveProblem(item)
	})

	println(lo.Sum(solvedProblems))
}
