package day_1

import "github.com/samber/lo"

func Part2(input string) {
	turns := parseDialTurns(input)
	dial := Dial{
		position: 50,
	}
	zeroPositionCount := 0
	lo.ForEach(turns, func(item DialTurn, _ int) {
		zeroPositionCount += turnDial(&dial, item, 0)
	})
	println(zeroPositionCount)
}
