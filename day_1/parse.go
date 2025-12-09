package day_1

import (
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func parseDialTurns(input string) []DialTurn {
	lines := strings.Split(input, "\n")
	return lo.Map(lines, func(item string, index int) DialTurn {
		direction := item[0:1]
		distance, _ := strconv.ParseInt(item[1:], 10, 64)
		return DialTurn{
			distance:  int(distance),
			direction: direction,
		}
	})
}
