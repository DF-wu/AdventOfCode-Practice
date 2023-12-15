package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func filereader() string {
	inputFile, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("file err")
		return ""
	}
	inputString := string(inputFile)
	return inputString
}

func main() {
	input := filereader()
	// p1(input)
	p1(input)
	// p2(input)
}

func p1(puzzle string){
	// deal with deserialization first 
	
}

