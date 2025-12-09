package day_1

import "github.com/samber/lo"

func Part1(input string) {
	turns := parseDialTurns(input)
	dial := Dial{
		position: 50,
	}
	zeroPositionCount := 0
	lo.ForEach(turns, func(item DialTurn, _ int) {
		if turnDial(&dial, item, 0) > 0 {
			zeroPositionCount += 1
		}
	})
	println(zeroPositionCount)
}
