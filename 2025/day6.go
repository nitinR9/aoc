package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day6() {
	data := readFile("inputs/day6.txt", "\r\n")
	fileData := make([]string, len(data))
	copy(fileData, data)

	numbers, maxNumLength := [][]int64{}, 0

	operators := strings.Split(strings.ReplaceAll(data[len(data)-1], " ", ""), "")
	data = data[:len(data)-1]

	for _, line := range data {
		num, numList := "", []int64{}
		
		for _, val := range line {
			if val != ' ' {
				num += string(val)
			} else {
				if len(num) > 0 {
					inVal, _ := strconv.ParseInt(num, 10, 64)
					numList = append(numList, inVal)
					if maxNumLength < len(num) {
						maxNumLength = len(num)
					}
					num = ""
				}
			}
		}

		if len(num) > 0 {
			inVal, _ := strconv.ParseInt(num, 10, 64)
			numList = append(numList, inVal)
			if maxNumLength < len(num) {
				maxNumLength = len(num)
			}
		}

		numbers = append(numbers, numList)
	}

	p1, k, rowMax := int64(0), 0, len(numbers)
	for k < len(operators) {
		result := int64(0)
		if operators[k] == "*" {
			result = 1
		}
		for row:= 0; row < rowMax; row++ {
			switch operators[k] {
			case "*":
				result *= numbers[row][k]
			case "+":
				result += numbers[row][k]
			}
		}
		p1 += result
		k++
	}

	p2, currentOps, result, numList := int64(0), rune(0), int64(0), []int64{}
	rowMax = len(fileData)-1

	for j:= 0; j < len(fileData[0]); j++ {
		if fileData[rowMax][j] != ' ' {
			if len(numList) > 0 {
				if currentOps == '*' {
					result = 1
				}
				for _, val := range numList {
					switch currentOps {
					case '*':
						result *= val
					case '+':
						result += val
					}
				}
				p2 += result
				numList = []int64{}
				result = 0
			}
			currentOps = rune(fileData[rowMax][j])
		}

		num := ""

		for row:= 0; row < rowMax; row++ {
			if data[row][j] != ' ' {
				num += string(fileData[row][j])
			}
		}

		if len(num) > 0 {
			inVal, _ := strconv.ParseInt(num, 10, 64)
			numList = append(numList, inVal)
		}
	}


	if len(numList) > 0 {
		if currentOps == '*' {
			result = 1
		}
		for _, val := range numList {
			switch currentOps {
			case '*':
				result *= val
			case '+':
				result += val
			}
		}
		p2 += result
	}


	fmt.Println("Day 6 Part 1:", p1)
	fmt.Println("Day 6 Part 2:", p2)
	fmt.Println("--------------------------")
}