package parser

import "strconv"

func ParseInt(s string) int {
	parsed, _ := strconv.ParseInt(s, 10, 64)
	return int(parsed)
}
