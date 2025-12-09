package day_3

import "github.com/samber/lo"

func findHighestJoltage(batteries []int) int {
	highestLeft := lo.Max(batteries[0 : len(batteries)-1])
	_, start, _ := lo.FindIndexOf(batteries[0:len(batteries)-1], func(item int) bool {
		return item == highestLeft
	})

	highestRight := lo.Max(batteries[start+1:])

	return highestLeft*10 + highestRight
}

func Part1(input string) {
	banks := parseBatteriesBanks(input)

	joltages := lo.Map(banks, func(batteries []int, index int) int {
		return findHighestJoltage(batteries)
	})

	println(lo.Sum(joltages))
}
