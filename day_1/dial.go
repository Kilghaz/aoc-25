package day_1

const (
	Left  string = "L"
	Right        = "R"
)

type DialTurn struct {
	direction string
	distance  int
}

type Dial struct {
	position int
}

var dialDirection = map[string]int{
	Left:  -1,
	Right: 1,
}

func turnDial(dial *Dial, turn DialTurn, zeroCount int) int {
	if turn.distance == 0 {
		return zeroCount
	}

	direction := dialDirection[turn.direction]
	dial.position += direction

	if dial.position > 99 {
		dial.position = 0
	}

	if dial.position < 0 {
		dial.position = 99
	}

	if dial.position == 0 {
		zeroCount += 1
	}

	return turnDial(dial, DialTurn{
		direction: turn.direction,
		distance:  turn.distance - 1,
	}, zeroCount)
}
