package day_3

import (
	"github.com/samber/lo"
)

func Part2(input string) {
	banks := parseBatteriesBanks(input)

	joltages := lo.Map(banks, func(batteries []int, index int) int {
		return findHighestJoltage(batteries, 12)
	})

	println(lo.Sum(joltages))
}
