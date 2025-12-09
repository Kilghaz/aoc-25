package main

import (
	"aoc25/day_1"
	"aoc25/day_2"
	"aoc25/day_3"
	"aoc25/day_4"
	"aoc25/day_5"
	"aoc25/day_6"
	"aoc25/day_7"
	"aoc25/day_8"
	"aoc25/day_9"
	"aoc25/io"
	"aoc25/parser"
	"fmt"
	"os"
	"time"
)

var advents = map[int][]func(input string){
	1: {day_1.Part1, day_1.Part2},
	2: {day_2.Part1, day_2.Part2},
	3: {day_3.Part1, day_3.Part2},
	4: {day_4.Part1, day_4.Part2},
	5: {day_5.Part1, day_5.Part2},
	6: {day_6.Part1, day_6.Part2},
	7: {day_7.Part1, day_7.Part2},
	8: {day_8.Part1, day_8.Part2},
	9: {day_9.Part1, day_9.Part2},
}

func main() {
	day := parser.ParseInt(os.Args[1])
	part := parser.ParseInt(os.Args[2]) - 1
	var input string
	if len(os.Args) == 3 || os.Args[3] != "test" {
		input = io.LoadInput(day)
	} else {
		input = io.LoadTestInput(day)
	}

	start := time.Now().UnixMicro()
	advents[day][part](input)
	end := time.Now().UnixMicro()

	println(fmt.Sprintf("\nRuntime:\n%dμs \n%fms", end-start, (float64)(end-start)/1_000))
}
