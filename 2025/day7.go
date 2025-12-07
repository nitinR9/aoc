package main

import (
	"fmt"
)

func travel(grid *[][]rune, row, col int, ans *int64) {
	rowMax, colMax := len(*grid), len((*grid)[0])

	if (col < 0 || col >= colMax) || (row < 0 || row >= rowMax) {
		return
	}

	char := (*grid)[row][col]
	if char == '|' {
		return
	}

	if char == '^' {
		(*ans)++
		travel(grid, row, col-1, ans)
		travel(grid, row, col+1, ans)
		return
	}

	(*grid)[row][col] = '|'
	travel(grid, row+1, col, ans)
}

func travel2(grid *[][]rune, row, col int, bgrid *[][]bool) {
	rowMax, colMax := len(*grid), len((*grid)[0])
	
	if (col < 0 || col >= colMax) || (row < 0 || row >= rowMax) {
		return
	}

	char := (*grid)[row][col]
	boolVal := (*bgrid)[row][col]

	if boolVal {
		return
	}

	if char == '^' {
		travel2(grid, row, col-1, bgrid)
		travel2(grid, row, col+1, bgrid)
		return
	}
	
	(*bgrid)[row][col] = true

	travel2(grid, row+1, col, bgrid)
}

func day7() {
	file := readFile("inputs/day7.txt", "\r\n")
	parsedFile := [][]rune{}
	r, c := 0,0

	for _, line := range file {
		parsedFile = append(parsedFile, []rune(line))
	}

	for row, line := range parsedFile {
		for col := range line {
			if parsedFile[row][col] == 'S' {
				r,c = row, col
				break
			}
		}
	}
	p1 := int64(0)
	travel(&parsedFile, r, c, &p1)

	p2 := int64(0)
	intGrid := make([][]int64, len(parsedFile))
	bGrid := make([][]bool, len(parsedFile))

	for i := range intGrid {
		intGrid[i] = make([]int64, len(parsedFile[0]))
		bGrid[i] = make([]bool, len(parsedFile[0]))
	}

	intGrid[r][c] = 1

	travel2(&parsedFile, r, c, &bGrid)

	rMax, cMax := len(intGrid), len(intGrid[0])

	for i:=1; i<rMax; i++ {
		for j:=0; j<cMax; j++ {
			if !bGrid[i][j] {
				continue
			}
			intGrid[i][j] = intGrid[i-1][j]
			if j-1 >=0 {
				if parsedFile[i][j-1] == '^' {
					intGrid[i][j] += intGrid[i-1][j-1]
				}
			}
			if j+1 < cMax {
				if parsedFile[i][j+1] == '^' {
					intGrid[i][j] += intGrid[i-1][j+1]
				}
			}
		}
	}
	for _, value := range intGrid[len(intGrid)-1] {
		p2 += value
	}
	fmt.Println("Day 6 part 1", p1)
	fmt.Println("Day 6 part 2", p2)
	fmt.Println("--------------------------")
}