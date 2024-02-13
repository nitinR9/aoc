package main

import (
	"2015/common"
	"fmt"
	"strings"
)

type pos struct {
	x int
	y int
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

func part1(input []string) {
	hashmap := map[string]int{}
	posmap := map[string]pos{
		"^": {x: -1, y: -1},
		"v": {x: 1, y: 1},
		">": {x: 1, y: 0},
		"<": {x: -1, y: 0},
	}
	curr_pos := newPos()
	count := 0
	for _, val := range input {
		string := ""

		if val == "^" {
			curr_pos.setPos(posmap["^"])
		} else if val == "v" {
			curr_pos.setPos(posmap["v"])
		} else if val == ">" {
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

	fmt.Println("Total houses delivered =", count)
}

func main() {
	common.GetFile("input.txt")
	result := strings.Split(common.Input, "")
	part1(result)
}
