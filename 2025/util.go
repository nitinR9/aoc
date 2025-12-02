package main

import (
	"os"
	"strconv"
	"strings"
)

func readFile(path string, delimiter string) []string{
	fileData, _ := os.ReadFile(path)
	parsedRes := strings.Split(string(fileData), delimiter)
	return parsedRes
}

func convToNum(x string) int{
	num, _ := strconv.Atoi(x)
	return num
}