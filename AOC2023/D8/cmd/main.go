package main

import (
	"fmt"
	"github.com/DF-wu/AdventOfCode-Practice/dfmisc"
	"strings"
)

type node struct {
	L string
	R string
}

func main() {
	input := dfmisc.Filereader("input.txt")
	//fmt.Println(p1(input))
	fmt.Println(p2(input))
	//fmt.Println(input)
}

/*
	LR

	11A = (11B, XXX)
	11B = (XXX, 11Z)
	11Z = (11B, XXX)
	22A = (22B, XXX)
	22B = (22C, 22C)
	22C = (22Z, 22Z)
	22Z = (22B, 22B)
	XXX = (XXX, XXX)
*/
// both simluate or math way is fine. I prefer math way. Find LCM
func p2(input string) string {
	recipe := strings.Split(input, "\r\n\r\n")[0]
	rawMatrix := strings.Split(input, "\r\n\r\n")[1]
	startingNodes := []string{}
	matrix := map[string]node{}
	for _, line := range strings.Split(rawMatrix, "\r\n") {
		var L, R, entry string
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		fmt.Sscanf(line, "%s = %s %s", &entry, &L, &R)
		matrix[entry] = node{L, R}
		if strings.HasSuffix(entry, "A") {
			startingNodes = append(startingNodes, entry)
		}
		fmt.Println(matrix[entry])
	}
	fmt.Println(recipe)
	// find LCM steps for each starting node
	LCMList := []int{}
	for _, startingNode := range startingNodes {
		pivot, ctr := 0, 0
		curr := startingNode
		for pivot <= len(recipe) {
			if pivot == len(recipe) {
				//reset
				pivot = 0
			}
			if strings.HasSuffix(curr, "Z") {
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
		LCMList = append(LCMList, ctr)
	}
	fmt.Println(LCMList)
	//LCM = ab / gcd(a,b)
	LCM := 0
	for i := 0; i < len(LCMList)-1; i++ {
		if i == 0 {
			LCM = LCMList[i] * LCMList[i+1] / gcd(LCMList[i], LCMList[i+1])
		} else {
			LCM = LCM * LCMList[i+1] / gcd(LCM, LCMList[i+1])
		}
	}
	fmt.Println(LCM)
	return ""
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
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
