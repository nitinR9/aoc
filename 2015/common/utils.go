package common

import (
	"fmt"
	"io"
	"os"
)

var Input = ""

func GetFile(path string) {
	fileP, err := os.Open(path)
	if err != nil {
		fmt.Println("Error occurred", err)
	}
	defer fileP.Close()

	content, err := io.ReadAll(fileP)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	Input = string(content)
}
