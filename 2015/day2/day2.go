package main

import (
	"2015/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func part1(result []string) {
	totalArea := 0.0

	for _, val := range result {
		string_dims := strings.Split(val, "x")
		var dim [3]float64
		for j, num := range string_dims {
			temp, err := strconv.ParseFloat(num, 64)
			if err != nil {
				fmt.Println("String conversion failed", err)
				return
			}
			dim[j] = temp
		}
		l, b, h := dim[0], dim[1], dim[2]
		lb := l * b
		bh := b * h
		hl := h * l

		min := lb
		if bh < min {
			min = bh
		}
		if hl < min {
			min = hl
		}
		calcResult := (2.0 * (lb + bh + hl)) + min
		totalArea += calcResult
	}

	fmt.Println("Total surface area:", int(totalArea))
}

func part2(result []string) {
	total := int64(0)

	for _, val := range result {
		res := strings.Split(val, "x")
		// fmt.Println(res)
		var dims [3]int64
		vol := int64(1)

		for i, val := range res {
			conv, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				fmt.Println("conversion error", err)
				return
			}
			dims[i] = conv
			vol *= conv
		}

		var s1 int64 = math.MaxInt64
		var s2 int64 = math.MaxInt64
		for _, side := range dims {
			if side < s1 {
				s2, s1 = s1, side
			} else if side < s2 {
				s2 = side
			}
		}
		ribbon := (2 * (s1 + s2)) + vol
		total += ribbon
	}

	fmt.Println("Total ribbon required", total)
}

func main() {
	result := common.Process("input.txt")
	part1(result)
	part2(result)
}
