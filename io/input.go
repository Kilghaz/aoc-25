package io

import (
	"fmt"
	"os"
	"strings"
)

func LoadInput(n int) string {
	dat, err := os.ReadFile(fmt.Sprintf("./inputs/input-%d.txt", n))
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(dat))
}

func LoadTestInput(n int) string {
	dat, err := os.ReadFile(fmt.Sprintf("./inputs/input-%d-test.txt", n))
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(dat))
}
