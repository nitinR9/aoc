package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Region struct {
	row int
	col int
	boxQuants []int
}

type Shape = []string

type Input struct {
	shapes []Shape
	regions []Region
}

func parseDay12(input []string) Input {
	var result Input
	blockNum, lineShape := 0, Shape{}
	for _, line := range input {
		if line == "" {
			result.shapes = append(result.shapes, lineShape)
			lineShape = Shape{} 
			blockNum++
		} else if strings.Contains(line, "#") {
			lineShape = append(lineShape, line)
		} else if strings.Contains(line, "x") {
			str := strings.Split(line, ":")
			dimStr := strings.Split(str[0], "x")
			row, _ := strconv.Atoi(dimStr[0]) 
			col, _ := strconv.Atoi(dimStr[1]) 

			boxQuants := strings.Split(str[1], " ")
			boxQuantInts := []int{}
			for _, q := range boxQuants {
				quant, _ := strconv.Atoi(q)
				boxQuantInts = append(boxQuantInts, quant)
			}
			result.regions = append(result.regions, Region{row: row, col: col, boxQuants: boxQuantInts})
		}
	}

	return result
}

func day12() {
	data := readFile("inputs/day12.txt", "\n")
	result := parseDay12(data)
	fmt.Println(result)
}