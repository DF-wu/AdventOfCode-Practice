package main

import (
	"fmt"
	"github.com/DF-wu/AdventOfCode-Practice/dfmisc"
	"strings"
)

func main() {
	input := dfmisc.Filereader("input.txt")
	fmt.Println(p1(input))
	//fmt.Println(p2(input))
	//fmt.Println(input)
}

/*
	RL

	AAA = (BBB, CCC)
	BBB = (DDD, EEE)
	CCC = (ZZZ, GGG)
	DDD = (DDD, DDD)
	EEE = (EEE, EEE)
	GGG = (GGG, GGG)
	ZZZ = (ZZZ, ZZZ)
*/

type node struct {
	L string
	R string
}

func p1(input string) string {
	//deserialise input.     Notice: this is windows line ending
	// use "\r\n\r\n" to split the input into two parts
	recipe := strings.Split(input, "\r\n\r\n")[0]
	rawMatrix := strings.Split(input, "\r\n\r\n")[1]
	matrix := map[string]node{}
	for _, line := range strings.Split(rawMatrix, "\r\n") {
		var L, R, entry string
		//format scanf can not read string but operate by comma. IDK why
		//replce to "". and remove () to make it work
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		fmt.Sscanf(line, "%s = %s %s", &entry, &L, &R)
		matrix[entry] = node{L, R}
	}
	fmt.Println(recipe)
	fmt.Println(matrix)

	// quiz specific start. end at "ZZZ"
	curr := "AAA"
	// pivot to point the curr char
	pivot, ctr := 0, 0
	for pivot <= len(recipe) {
		if pivot == len(recipe) {
			//reset
			pivot = 0
		}
		if curr == "ZZZ" {
			// found
			break
		}

		switch string(recipe[pivot]) {
		case "R":
			curr = matrix[curr].R
		case "L":
			curr = matrix[curr].L
		default:
			fmt.Println("error")
		}
		ctr++
		pivot++

	}
	fmt.Println(ctr)

	return ""
}
