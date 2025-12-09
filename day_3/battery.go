package day_3

import (
	"aoc25/math"

	"github.com/samber/lo"
)

func findHighestJoltage(batteries []int, batteryCount int) int {
	if batteryCount == 0 {
		return 0
	}

	batteriesToSearch := batteries[:len(batteries)-batteryCount+1]
	maxJoltageBattery := lo.Max(batteriesToSearch)
	_, batteryIndex, _ := lo.FindIndexOf(batteriesToSearch, func(item int) bool {
		return item == maxJoltageBattery
	})

	return maxJoltageBattery*math.Pow(10, batteryCount-1) + findHighestJoltage(batteries[batteryIndex+1:], batteryCount-1)
}
