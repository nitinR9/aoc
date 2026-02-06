package main

import (
	"fmt"
	"strings"
)

func countDevices(deviceMap *map[string][]string, start, end string, memo map[string]int64) int64{
	key := start + "-" + end
	if val, exists := memo[key]; exists {
		return val
	}
	if start == end {
		memo[key] = 1
		return 1
	}
	total := int64(0)
	for _, device := range (*deviceMap)[start] {
		total += countDevices(deviceMap, device, end, memo)
	}
	memo[key] = total
	return total
}

func day11part1(deviceMap *map[string][]string) int64{
	count := countDevices(deviceMap, "you", "out", make(map[string]int64))
	return count
}

func day11part2(deviceMap *map[string][]string) int64{
	mapCount := make(map[string]int64)
	return (countDevices(deviceMap, "svr", "fft", mapCount) * countDevices(deviceMap, "fft", "dac", mapCount) * countDevices(deviceMap, "dac", "out", mapCount)) + (countDevices(deviceMap, "svr", "dac", mapCount) * countDevices(deviceMap, "dac", "fft", mapCount) * countDevices(deviceMap, "fft", "out", mapCount))
}

func parseInputDay11(input []string) map[string][]string {
	result := make(map[string][]string)
	for _, line := range input {
		splitStr := strings.Split(line, ":")
		result[splitStr[0]] = strings.Split(strings.TrimSpace(splitStr[1]), " ")
	}

	return result
}

func day11() {
	data := readFile("inputs/day11.txt", "\n")
	deviceMap := parseInputDay11(data)
	fmt.Println("Day 11 Part 1:", day11part1(&deviceMap))
	fmt.Println("Day 11 Part 2:", day11part2(&deviceMap))
}