package day_6

import (
	"aoc25/parser"
	"aoc25/slice"
	"aoc25/str"
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
		return MathProblem{numbers: problemNumbers, operator: []rune(operator)[0]}
	})
}

func parseCephalopodWorksheet(input string) []MathProblem {
	lines := strings.Split(input, "\n")
	maxLineLength := lo.Max(lo.Map(lines, func(line string, _ int) int {
		return len(line)
	}))
	lines = lo.Map(lines, func(line string, _ int) string {
		return str.PadRight(line, " ", maxLineLength)
	})
	operatorLine := []rune(slice.Last(lines))
	numbersGrid := lo.Map(lines[:len(lines)-1], func(line string, index int) []rune {
		return []rune(line)
	})

	var mathProblems []MathProblem
	currentMathProblem := MathProblem{}

	readNumberColumn := func(index int) int {
		column := lo.Map(numbersGrid, func(row []rune, _ int) rune {
			return row[index]
		})
		column = lo.Filter(column, func(item rune, index int) bool {
			return item != ' '
		})
		return parser.ParseInt(string(column))
	}

	pushMathProblem := func() {
		mathProblems = append(mathProblems, currentMathProblem)
		currentMathProblem = MathProblem{}
	}

	for i := 0; i < len(operatorLine); i++ {
		columnNumber := readNumberColumn(i)
		if columnNumber == 0 {
			pushMathProblem()
			continue
		}

		if operatorLine[i] != ' ' {
			currentMathProblem.operator = operatorLine[i]
		}

		currentMathProblem.numbers = append(currentMathProblem.numbers, readNumberColumn(i))
	}
	pushMathProblem()

	return mathProblems
}
