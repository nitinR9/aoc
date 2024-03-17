package main

import (
	"2015/common"
	"fmt"
	"strconv"
)

func ProcessLength(s string) int {
	s, _ = strconv.Unquote(s)
	return len(s)
}

func Part1(strings []string) {
	total := 0
	for _, s := range strings {
		total += len(s) - ProcessLength(s)
	}

	fmt.Println("Part1:", total)
}

func Part2(strings []string) {
	total := 0
	for _, s := range strings {
		total += len(strconv.Quote(s)) - len(s)
	}

	fmt.Println("Part2:", total)
}

func main() {
	input := common.Process("input.txt")
	Part1(input)
	Part2(input)
}
