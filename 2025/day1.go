package main

import (
	"fmt"
	"strconv"
)

type Instruction struct{
	direction rune
	postitions int
}

func day1(){
	parsedRes := readFile("inputs/day1.txt", "\r\n")
	result := []Instruction{}
	for _, v := range parsedRes {
		pos, _ := strconv.Atoi(v[1:])
		result = append(result, Instruction{ direction: rune(v[0]), postitions: pos })
	}

	pos, p1, p2 := 50, 0, 0
	for _, v := range result {
		for i:=0 ; i<v.postitions; i++{
			if v.direction == 'L' {
				pos = (pos - 1 + 100)%100
			} else {
				pos = (pos + 1)%100
			}
			if pos == 0 {
				p2++
			}
		}
		
		if pos == 0 {
			p1++
		}
	}

	fmt.Println("Day 1 P1:", p1)
	fmt.Println("Day 1 P2:", p2)
	fmt.Println("--------------------------")
}