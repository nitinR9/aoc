package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func combineNums(nums []int) int64{
	val, k := int64(0), int64(len(nums))-1

	for _, num := range nums {
		val += int64(math.Pow(10, float64(k))) * int64(num)
		k--
	}

	return val
}

func conv(nums ...int) int64 {
    var b strings.Builder
    for _, n := range nums {
        b.WriteString(strconv.Itoa(n))
    }

    result, _ := strconv.ParseInt(b.String(), 10, 64)
    return result
}

func getLargestNum(b []int) int64 {
	digits := []int{}
	skips := len(b) - 12
	currentIndex := 0
	for len(digits) < 12 {
		next_digit := slices.Max(b[currentIndex : currentIndex + skips + 1])
		for b[currentIndex] != next_digit {
			currentIndex++
			skips--
		}
		digits = append(digits, next_digit)
		currentIndex++
	}
	return conv(digits...)
}

func day3() {
	data := readFile("inputs/day3.txt", "\r\n")
	batteries := [][]int{}
	for _,v := range data {
		splitBatteries := strings.Split(v, "")
		intValues := make([]int, len(splitBatteries))
		for i, val := range splitBatteries {
			jolts, _ := strconv.Atoi(val)
			intValues[i] = jolts
		}
		batteries = append(batteries, intValues)
	}

	p1,p2,bankLength := int64(0),int64(0),len(batteries[0])-1

	for _, bank := range batteries {
		max2digits := [2]int{}
		for i, jolts := range bank {
			if jolts > max2digits[0] && i != bankLength {
				max2digits = [2]int{jolts, 0}
			} else if jolts > max2digits[1] {
				max2digits[1] = jolts
			}
		}
		combinedJolts := combineNums(max2digits[:])
		p1 += combinedJolts
	}

	fmt.Println("Day 3 P1:", p1)

	for _, bank := range batteries {
		num := getLargestNum(bank)
		p2 += num
	}
	
	fmt.Println("Day 3 P2:", p2)
	fmt.Println("--------------------------")
}