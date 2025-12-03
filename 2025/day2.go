package main

import (
	"fmt"
	"strconv"
	"strings"
)

func checkIdP1(id string) bool {
	mid := len(id)/2
	return strings.Compare(id[:mid], id[mid:]) == 0
}

func checkIdP2(id string) bool {
	mid := len(id)/2
	
	for mid >= 1 {
		substr := id[:mid]
		count := strings.Count(id, substr)
		if  len(id) == (count * mid) {
			return true
		}
		mid--
	}

	return false
}

func day2() {
	data := readFile("inputs/day2.txt", ",")
	parsedRes := [][2]int{}
	for _, v := range data {
		removeNewLines := strings.ReplaceAll(v, "\r\n", "")
		splitVal := strings.Split(removeNewLines, "-")
		parsedRes = append(parsedRes, [2]int{convToNum(splitVal[0]), convToNum(splitVal[1])})
	}

	p1, p2 := 0, 0

	for _, v := range parsedRes {
		first, last := v[0], v[1]

		for first <= last {
			if checkIdP1(strconv.Itoa(first)) {
				p1 += first
			}

			if checkIdP2(strconv.Itoa(first)) {
				p2 += first
			}
			first++
		}
	}

	fmt.Println("Day 2 P1:", p1)
	fmt.Println("Day 2 P2:", p2)
	fmt.Println("--------------------------")
}