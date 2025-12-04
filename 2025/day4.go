package main

import (
	"fmt"
	"strings"
)

func hasAdjacent(grid *[][]string, i int, j int) bool {
	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}
	
	countRolls := 0
	
	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if ni >= 0 && ni < len(*grid) && nj >= 0 && nj < len((*grid)[0]) {
			if (*grid)[ni][nj] == "@" {
				countRolls++
			}
		}
	}
	return countRolls < 4
}

func day4() {
	data := readFile("inputs/day4.txt", "\r\n")
	grid := [][]string{}
	
	for _, line := range data {
		grid = append(grid, strings.Split(line, ""))
	}
	
	p1 := 0
	p2 := 0
	for i, line := range grid {
		for j, value := range line {
			if value == "@" && hasAdjacent(&grid, i, j) {
				p1++
			}
		}
	}
	
	gridChanged := true
	for gridChanged {
		gridChanged = false
		for i, line := range grid {
			for j, value := range line {
				if value == "@" && hasAdjacent(&grid, i, j) {
					grid[i][j] = "."
					p2++
					gridChanged = true
				}
			}
		}
	}
	
	fmt.Println("Day 4 P1:", p1)
	fmt.Println("Day 4 P2:", p2)
	fmt.Println("--------------------------")
}