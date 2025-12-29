package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type Vec2d struct {
	x int
	y int
}

type BitSet struct {
	rows int
	cols int
	grid []uint64
}

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

func NewBitSet(rows, cols int) *BitSet {
	totalBits := rows * cols
	totalSetSize := (totalBits+63) / 64
	return &BitSet{
		rows: rows,
		cols: cols,
		grid: make([]uint64, totalSetSize),
	}
}

func (b *BitSet) hasTile(r, c int) bool {
	idx := (r * b.cols) + c
	return (b.grid[idx >> 6] & uint64(1 << (idx & 63))) != 0
}

func (b *BitSet) addTilesToGrid(p1, p2 Vec2d) {
	start, end := 0, 0
	if p1.x == p2.x {
		start, end = p1.y, p2.y
		if start > end {
			start, end = end, start
		}
		for y := start; y <= end; y++ {
			idx := (y * b.cols) + p1.x
			b.grid[idx >> 6] |= uint64(1 << (idx & 63))
		}
	} else if p1.y == p2.y {
		start, end = p1.x, p2.x
		if start > end {
			start, end = end, start
		}
		for x := start; x <= end; x++ {
			idx := (p1.y * b.cols) + x
			b.grid[idx >> 6] |= uint64(1 << (idx & 63))
		}
	}
}

func (b *BitSet) processRow(row int) {
	cS, cE := 0, b.cols-1
	found := false
	for cS <= cE {
		fS, fE := b.hasTile(row,cS), b.hasTile(row,cE)
		if !found {
			if fS && fE {
				found = true
			} else if !fS {
				cS++
			} else if !fE {
				cE--
			}
		} else {
			idx := 0
			if cS != cE {
				idx = (row * b.cols) + cS
				b.grid[idx >> 6] |= uint64(1 << (idx & 63))
				idx = (row * b.cols) + cE
				b.grid[idx >> 6] |= uint64(1 << (idx & 63))
			} else {
				idx = (row * b.cols) + cS
				b.grid[idx >> 6] |= uint64(1 << (idx & 63))
			}
			cS++
			cE--
		}
	}
}

func (b *BitSet) fillWithTilesMultithreaded() {
	var wg sync.WaitGroup
	rS, rE := 0, b.rows-1

	for rS <= rE {
		if rS == rE {
			wg.Add(1)
			go func(row int) {
				defer wg.Done()
				b.processRow(row)
			}(rS)
		} else {
			wg.Add(2)
			go func(row int) {
				defer wg.Done()
				b.processRow(row)
			}(rS)
			go func(row int) {
				defer wg.Done()
				b.processRow(row)
			}(rE)
		}
		rS++
		rE--
	}
	wg.Wait()
}

func (b *BitSet) fillWithTilesSingleThreaded() {
	rS, rE := 0, b.rows-1
	for rS <= rE {
		b.processRow(rS)
		rS++
	}	
}

func (b *BitSet) isOutsideBorders(p1, p2 Vec2d) bool {
	if p1.x == p2.x {
		start, end := p1.y, p2.y
		if start > end {
			start, end = end, start
		}
		for y := start; y <= end; y++ {
			if !b.hasTile(y, p1.x) {
				return true
			}
		}
	} else if p1.y == p2.y {
		start, end := p1.x, p2.x
		if start > end {
			start, end = end, start
		}
		for x := start; x <= end; x++ {
			if !b.hasTile(p1.y, x) {
				return true
			}
		}
	}
	return false
}

func (b *BitSet) isRecOutside(p1, p2 Vec2d) bool {
	if p1.x == p2.x || p1.y == p2.y {
		return b.isOutsideBorders(p1, p2)
	}

	points := []Vec2d{p1, {p1.x, p2.y}, p2, {p2.x, p1.y}, p1}

	// check if corners are outside grid of tiles (speed boost)
	if !b.hasTile(p2.y, p1.x) || !b.hasTile(p1.y, p2.x) {
		return true
	}

	// check all four borders if not outside grid of tiles
	for i:=0; i<len(points)-1; i++ {
		if b.isOutsideBorders(points[i], points[i+1]) {
			return true
		}
	}

	return false
}

func (b *BitSet) getMaxAreaRectangle(grid *[]Vec2d) int64 {
	maxArea := int64(0)

	for i := 0; i < len(*grid)-1; i++ {
		for j := 0; j < len(*grid); j++ {
			g1, g2 := (*grid)[i], (*grid)[j]
			if !b.isRecOutside(g1, g2) {
				a := Area(g1, g2)
				if a > maxArea {
					maxArea = a
				}
			}
		}
	}

	return maxArea
}

func part2day9Slow(grid []Vec2d, dim Vec2d) int64 {
	bs := NewBitSet(dim.y+1, dim.x+1)

	for i:=0; i<len(grid)-1; i++ {
		p1, p2 := grid[i], grid[i+1]
		bs.addTilesToGrid(p1, p2)
	}

	bs.addTilesToGrid(grid[len(grid)-1], grid[0])

	bs.fillWithTilesSingleThreaded()

	maxArea := int64(0)
	for i := 0; i < len(grid)-1; i++ {
		for j := 0; j < len(grid); j++ {
			if !bs.isRecOutside(grid[i], grid[j]) {
				a := Area(grid[i], grid[j])
				if a > maxArea {
					maxArea = a
				}
			}
		}
	}

	return maxArea
}

func part2day9Fast(grid []Vec2d, dMax Vec2d) int64 {

	bs := NewBitSet(dMax.y+1, dMax.x+1)

	for i:=0; i<len(grid)-1; i++ {
		p1, p2 := grid[i], grid[i+1]
		bs.addTilesToGrid(p1, p2)
	}

	bs.addTilesToGrid(grid[len(grid)-1], grid[0])

	bs.fillWithTilesMultithreaded()

	maxArea := bs.getMaxAreaRectangle(&grid)
	
	return maxArea
}

func part1day9(grid []Vec2d) (int64, Vec2d) {
	max, pMax := int64(0), Vec2d{0,0}
	for i := 0; i < len(grid)-1; i++ {
		for j := i + 1; j < len(grid); j++ {
			result := Area(grid[i], grid[j])
			if result > max {
				max = result
			}
			
			pMax = Vec2d{x: Max(grid[i].x, grid[j].x, pMax.x), y: Max(grid[i].y, grid[j].y, pMax.y)}
		}
	}
	return max, pMax
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
	
	p1, pMax := part1day9(grid)
	fmt.Println("Day 9 part 1:", p1)

	p2 := part2day9Fast(grid, pMax)
	fmt.Println("Day 9 part 2:", p2)
	fmt.Println("--------------------------")
}