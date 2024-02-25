package main

import (
	"2015/common"
	"fmt"
)

func main() {
	var floor int64 = 0
	var input = common.GetFile("input.txt")
	var i = 0
	for ; i < len(input); i++ {
		if input[i] == '(' {
			floor++
		}
		if input[i] == ')' {
			floor--
		}
		// extra code for part 2
		if floor == -1 {
			break
		}
	}

	fmt.Println("Floor reached", floor)
	fmt.Println("Reached basement at position", i+1)
}
