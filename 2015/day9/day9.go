package main

import (
	"2015/common"
	"fmt"
	"strconv"
	"strings"
)

type NextLocation struct {
	next     string
	distance int
}

func Part1(input []string) {
	locMap := map[string]NextLocation{}
	for _, line := range input {
		split := strings.Split(line, " ")
		from, to, distance := split[0], split[2], split[4]

		if val, err := strconv.ParseInt(distance, 10, 32); err != nil {
			fmt.Println("Error occurred when converting")
			return
		} else {
			if _, ok := locMap[from]; ok {
				locMap[to] = NextLocation{
					next:     from,
					distance: int(val),
				}
			} else {
				locMap[from] = NextLocation{
					next: to,
					distance: int(val),
				}
			}
		}
	}

	fmt.Println(locMap)
}

func main() {
	input := common.Process("test.txt")
	Part1(input)
}
