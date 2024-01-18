package dfmisc

import (
	"fmt"
	"os"
)

func filereader(path string) string {
	inputFile, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("file err")
		return ""
	}
	inputString := string(inputFile)
	return inputString
}
