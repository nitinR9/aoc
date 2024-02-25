package common

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
)

type OperatorFunction[T any] func(string) T

func GetFile(path string) string {
	fileP, err := os.Open(path)
	if err != nil {
		fmt.Println("Error occurred", err)
	}
	defer fileP.Close()

	content, err := io.ReadAll(fileP)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return string(content)
}

func Process(file string) []string {
	input := GetFile(file)
	seprator := ""
	if runtime.GOOS == "windows" {
		seprator = "\r\n"
	} else {
		seprator = "\n"
	}

	return strings.Split(input, seprator)
}

func Map(arr []string, fn OperatorFunction[uint16]) []uint16 {
	result := []uint16{}
	for _, val := range arr {
		result = append(result, fn(val))
	}
	return result
}
