package day_6

import "github.com/samber/lo"

func Part2(input string) {
	problems := parseCephalopodWorksheet(input)

	solvedProblems := lo.Map(problems, func(item MathProblem, _ int) int {
		return solveProblem(item)
	})

	println(lo.Sum(solvedProblems))
}
