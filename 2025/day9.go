package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Vec2d struct {
	x int
	y int
}

type ByteSlice []byte

func Area(p1 Vec2d, p2 Vec2d) int64 {
	width := int64(math.Abs(float64(p2.x-p1.x))) + 1
	height := int64(math.Abs(float64(p2.y-p1.y))) + 1

	return width * height
}

func Max(nums ...int) int {
	return slices.Max(nums)
}

func Min(nums ...int) int {
	return slices.Min(nums)
}

func (b *ByteSlice) setPixel(width, r, c int) {
	idx := (r * width) + c
	byteIdx := idx/8
	bitIdx := idx % 8
	(*b)[byteIdx] |= 1 << bitIdx
}

func (b *ByteSlice) hasPixel(width, r, c int) bool {
	idx := (r * width) + c
	byteIdx := idx/8
	bitIdx := idx % 8

	return (*b)[byteIdx] & (1 << bitIdx) != 0
}

func (b *ByteSlice) createWall(width int, p1, p2 Vec2d) {
	if p1.x == p2.x {
		start, end := 0, 0
		if p1.y < p2.y {
			start, end = p1.y, p2.y
		} else {
			start, end = p2.y, p1.y
		}
		for j:=start; j<=end; j++ {
			b.setPixel(width, j, p1.x)
		}
	} else if p2.y == p1.y {
		start, end := 0, 0
		if p1.x < p2.x {
			start, end = p1.x, p2.x
		} else {
			start, end = p2.x, p1.x
		}
		for i:=start; i<=end; i++ {
			b.setPixel(width, p1.y, i)
		}
	}
}

func (b *ByteSlice) fillWithWalls(width, height int) {
	for r:=0; r<height; r++ {
		foundIdx := -1
		for c := 0; c < width; c++ {
			if b.hasPixel(width, r,c) {
				if foundIdx != -1 {
					if (c - foundIdx > 1) {
						b.createWall(width, Vec2d{foundIdx+1,r}, Vec2d{c-1, r})
					}
				}
				foundIdx = c
			}
		}
	}
}

func (b *ByteSlice) isOutsideBorders(width int, p1, p2 Vec2d) bool {
	if p1.x == p2.x {
		start, end := 0, 0
		if p1.y < p2.y {
			start, end = p1.y, p2.y
		} else {
			start, end = p2.y, p1.y
		}
		for j:=start; j<=end; j++ {
			if !b.hasPixel(width, j, p1.x) {
				return true
			}
		}
	} else if p2.y == p1.y {
		start, end := 0, 0
		if p1.x < p2.x {
			start, end = p1.x, p2.x
		} else {
			start, end = p2.x, p1.x
		}
		for i:=start; i<=end; i++ {
			if !b.hasPixel(width, p1.y, i) {
				return true
			}
		}
	}
	return false
}

func (b *ByteSlice) checkIfRecOutside(width int, p1, p2 Vec2d) bool {
	
	if p1.x == p1.y || p1.y == p2.y {
		return b.isOutsideBorders(width, p1, p2)
	}

	p3, p4 := Vec2d{p1.x, p2.y}, Vec2d{p2.x, p1.y}

	points := []Vec2d{p1,p3,p2,p4,p1}

	for i :=0; i<len(points)-1; i++ {
		if b.isOutsideBorders(width, points[i], points[i+1]) {
			return true
		}
	}

	return false
}

func part2day9(grid []Vec2d, dim int) int64 {
	totalBits := dim * dim
	totalBytes := (totalBits + 7) / 8
	byteSet := make(ByteSlice, totalBytes)

	for i := 0; i<len(grid)-1; i++ {
		p1, p2 := grid[i], grid[i+1]
		byteSet.createWall(dim, p1, p2)
	}
	byteSet.createWall(dim, grid[len(grid)-1], grid[0])

	byteSet.fillWithWalls(dim, dim)

	maxArea := int64(0)
	
	for i := 0; i<len(grid)-1; i++ {
		for j := 0; j < len(grid); j++ {
			if !byteSet.checkIfRecOutside(dim, grid[i], grid[j]){
				result := Area(grid[i], grid[j])
				if result > maxArea {
					maxArea = result
				}
			}
		}
	}

	return maxArea
}

func part1day9(grid []Vec2d) (int64, int) {
	max, maxDim := int64(0), 0
	for i := 0; i < len(grid)-1; i++ {
		for j := i + 1; j < len(grid); j++ {
			result := Area(grid[i], grid[j])
			if result > max {
				max = result
			}

			maxDim = Max(maxDim, grid[i].x, grid[i].y, grid[j].x, grid[j].y)
		}
	}
	return max, maxDim
}

func day9() {
	data := readFile("inputs/day9.txt", "\r\n")
	grid := []Vec2d{}

	for _, line := range data {
		split := strings.Split(line, ",")
		xpos, _ := strconv.Atoi(split[0])
		ypos, _ := strconv.Atoi(split[1])
		grid = append(grid, Vec2d{x: xpos, y: ypos})
	}


	p1, charDim := part1day9(grid)
	fmt.Println("Day 9 part 1:", p1)
	p2 := part2day9(grid, charDim+1)
	fmt.Println("Day 9 part 2:", p2)
	fmt.Println("--------------------------")
}
