package main

import (
	"2015/common"
	"fmt"
)

type pos struct {
	x int
	y int
}

var posmap = map[string]pos{
	"^": {x: -1, y: -1},
	"v": {x: 1, y: 1},
	">": {x: 1, y: 0},
	"<": {x: -1, y: 0},
}

func newPos() pos {
	p := pos{
		x: 0,
		y: 0,
	}
	return p
}

func (p *pos) setPos(dir pos) {
	p.x += dir.x
	p.y += dir.y
}

func part1(input string) {
	hashmap := map[string]int{}

	curr_pos := newPos()
	count := 0
	for _, val := range input {
		string := ""

		if val == '^' {
			curr_pos.setPos(posmap["^"])
		} else if val == 'v' {
			curr_pos.setPos(posmap["v"])
		} else if val == '>' {
			curr_pos.setPos(posmap[">"])
		} else {
			curr_pos.setPos(posmap["<"])
		}
		string = fmt.Sprintf("%+v", curr_pos)
		if _, ok := hashmap[string]; !ok {
			hashmap[string] = 1
			count++
		}
	}

	fmt.Println("Part1:", count)
}

func (p *pos) move(h map[string]int, char rune) int {
	if char == '^' {
		p.setPos(posmap["^"])
	} else if char == 'v' {
		p.setPos(posmap["v"])
	} else if char == '>' {
		p.setPos(posmap[">"])
	} else {
		p.setPos(posmap["<"])
	}
	string := fmt.Sprintf("%+v", p)
	if _, ok := h[string]; !ok {
		h[string] = 1
		return 1
	}
	return 0
}

func part2(input string) {
	hash := map[string]int{}

	santa := newPos()
	robot := newPos()
	count := 0
	for index, char := range input {
		if index%2 == 0 {
			count += santa.move(hash, char)
		} else {
			count += robot.move(hash, char)
		}
	}

	fmt.Println("Part2:", count)
}

func main() {
	input := common.GetFile("test.txt")
	part1(input)
	part2(input)
}
