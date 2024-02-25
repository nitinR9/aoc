package main

import (
	"2015/common"
	"fmt"
	"strings"
)

func Check1(s string) bool {
	vowel, double, isOk := 0, 0, false

	if strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "xy") || strings.Contains(s, "pq") {
		return false
	} else {
		isOk = true
	}

	for i := 0; i < len(s)-1; i++ {
		l, next := string(s[i]), string(s[i+1])

		if strings.Contains("aeiou", l) {
			vowel++
		}
		if l == next {
			double++
		}
	}
	if strings.Contains("aeiou", string(s[len(s)-1])) {
		vowel++
	}
	return vowel >= 3 && double >= 1 && isOk
}

func Part1(input string, num int) {
	count := 0
	strings := strings.Split(input, "\n")

	for _, string := range strings {
		if Check1(string) {
			count++
		}
	}

	fmt.Printf("Part%d: %d\n", num, count)
}

func CheckRepeat(s string, start int) bool {
	for i := 1; i < len(s)-1; i++ {
		l, r := i-1, i+1

		if s[l] == s[r] {
			return true
		}
	}

	return false
}

func Check2(s string) bool {
	result := [2]int{0, 0}
	hash := map[string][2]int{}
	rule1 := false
	rule2 := false

	for i := 0; i < len(s)-1; i++ {
		string := fmt.Sprintf("%c%c", s[i], s[i+1])
		if val, ok := hash[string]; ok {
			if i > val[0]+1 {
				result = [2]int{val[0], i + 1}
				rule1 = true
				break
			}
		} else {
			hash[string] = [2]int{i, 0}
		}
	}

	_ = result

	if !rule1 {
		return false
	}

	for i := 1; i < len(s)-1; i++ {
		res := CheckRepeat(s, i)
		if res {
			rule2 = true
			break
		}
	}

	return rule1 && rule2
}

func Part2(input string, num int) {
	count := 0
	strings := strings.Split(input, "\n")

	for _, value := range strings {
		if Check2(value) {
			count++
		}
	}

	fmt.Printf("Part%d: %d\n", num, count)
}

func main() {
	input := common.GetFile("input.txt")
	Part1(input, 1)
	Part2(input, 2)
}
