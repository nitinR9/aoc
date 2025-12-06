package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int64
	end   int64
}

func parseData(data []string) ([]Range, []int64) {
	ranges := []Range{}
	nums := []int64{}

	for _, line := range data {
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			start, _ := strconv.ParseInt(parts[0], 10, 64)
			end, _ := strconv.ParseInt(parts[1], 10, 64)
			ranges = append(ranges, Range{start: start, end: end})
		} else {
			num, _ := strconv.ParseInt(line, 10, 64)
			nums = append(nums, num)
		}
	}

	return ranges, nums
}

func getMergedRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return []Range{}
	}
	
	sort.Slice(ranges, func(i, j int) bool {
        return ranges[i].start < ranges[j].start
    })

	merged := []Range{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		last := &merged[len(merged)-1]
		current := ranges[i]
		if current.start <= last.end+1 {
			if current.end > last.end {
				last.end = current.end
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}	

func day5() {
	data := readFile("inputs/day5.txt", "\r\n")
	index := slices.Index(data, "")
	data = slices.Delete(data, index, index+1)
	ranges, nums := parseData(data)

	p1 := 0
	
	for _, num := range nums {
		for _, r := range ranges {
			if num >= r.start && num <= r.end {
				p1++
				break
			}
		}
	}

	filteredRanges := getMergedRanges(ranges)

	p2 := int64(0)
	
	for _, rg := range filteredRanges {
		p2 = p2 + (rg.end - rg.start + 1)
	}

	fmt.Println("Day 5 Part 1:", p1)
	fmt.Println("Day 5 Part 2:", p2)
	fmt.Println("--------------------------")
}