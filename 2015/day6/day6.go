package main

import (
	"2015/common"
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type Instruction struct {
	state  int
	bounds []int
}

func Move1(arr *[1000][1000]int, state int, bounds []int) {
	x1, y1, x2, y2 := bounds[0], bounds[1], bounds[2], bounds[3]
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if state != -1 {
				arr[i][j] = state
			} else {
				arr[i][j] ^= 1
			}
		}
	}
}

func Move2(arr *[1000][1000]int, state int, bounds []int) {
	x1, y1, x2, y2 := bounds[0], bounds[1], bounds[2], bounds[3]
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			switch state {
			case 0:
				if arr[i][j] >= 1 {
					arr[i][j]--
				}
			case 1:
				arr[i][j]++
			default:
				arr[i][j] += 2
			}
		}
	}
}

func ConvertInt(s []string) []int {
	var res []int
	for _, str := range s {
		if val, err := strconv.ParseInt(str, 10, 64); err != nil {
			fmt.Println("Error converting int", err)
		} else {
			res = append(res, int(val))
		}
	}

	return res
}

func ProcessInput(s string) []Instruction {
	var ins []string
	if runtime.GOOS == "windows" {
		ins = strings.Split(s, "\r\n")
	} else {
		ins = strings.Split(s, "\n")
	}

	regex := regexp.MustCompile(`\d+`)
	result := []Instruction{}

	for _, val := range ins {
		bounds := ConvertInt(regex.FindAllString(val, -1))
		state := -1
		if strings.Contains(val, "turn off") {
			state = 0
		} else if strings.Contains(val, "turn on") {
			state = 1
		}

		result = append(result, Instruction{
			state,
			bounds,
		})
	}

	return result
}

func Part1(input string) {
	instructions := ProcessInput(input)
	var grid [1000][1000]int
	ans := 0

	for _, move := range instructions {
		Move1(&grid, move.state, move.bounds)
	}

	for _, val1 := range grid {
		for _, val2 := range val1 {
			if val2 == 1 {
				ans++
			}
		}
	}

	fmt.Println("Part1:", ans)
}

func Part2(input string) {
	instructions := ProcessInput(input)
	var grid [1000][1000]int
	ans := 0

	for _, move := range instructions {
		Move2(&grid, move.state, move.bounds)
	}

	for _, val1 := range grid {
		for _, val2 := range val1 {
			ans += val2
		}
	}

	fmt.Println("Part2:", ans)
}

func main() {
	input := common.GetFile("input.txt")
	Part1(input)
	Part2(input)
}
