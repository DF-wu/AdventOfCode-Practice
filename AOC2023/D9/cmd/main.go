/*
	Day 9
	Author: DF
	Date: 2024-01-25
*/

package main

import (
	"fmt"
	"github.com/DF-wu/AdventOfCode-Practice/dfmisc"
	"strconv"
	"strings"
)

func main() {
	input := dfmisc.Filereader("input.txt")
	fmt.Println(p1(input))
	fmt.Println(p2(input))
	//fmt.Println(input)
}

// similar to p1, but minus to get the target number backward
func p2(input string) int {
	ans := 0
	var currPyramid [][]int
	for _, quiz := range readQuiz(input) {
		currPyramid = readAsPyramid(quiz)

		// init ctr to count level
		levelCtr := len(currPyramid) - 1
		// init : copy first value as target number
		targetNumber := currPyramid[levelCtr][0]
		// append to the first position
		// the prepend method from https://stackoverflow.com/questions/17555857/go-unpacking-array-as-arguments
		currPyramid[levelCtr] = append([]int{targetNumber}, currPyramid[levelCtr]...)
		levelCtr--
		for levelCtr >= 0 {
			targetNumber = currPyramid[levelCtr][0] - currPyramid[levelCtr+1][0]
			currPyramid[levelCtr] = append([]int{targetNumber}, currPyramid[levelCtr]...)
			levelCtr--
		}

		ans += currPyramid[0][0]
	}
	return ans
}

/*
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
*/
func p1(input string) int {
	ans := 0
	for _, quiz := range readQuiz(input) {
		currPyramid := readAsPyramid(quiz)

		// start to extrapolate
		// in each level, the target number is the sum of the last number in this level and the last number in the last level

		// from the bottom level
		levelCtr := len(currPyramid) - 1
		// the bottom level is clear to get the target number. just copy the first number as target number
		currPyramid[levelCtr] = append(currPyramid[levelCtr], currPyramid[levelCtr][0])
		levelCtr--
		for levelCtr >= 0 {
			targetNumber := currPyramid[levelCtr][len(currPyramid[levelCtr])-1] + currPyramid[levelCtr+1][len(currPyramid[levelCtr+1])-1]
			currPyramid[levelCtr] = append(currPyramid[levelCtr], targetNumber)
			levelCtr--
		}
		ans += currPyramid[0][len(currPyramid[0])-1]
		fmt.Println(currPyramid)
	}

	return ans
}

func readQuiz(input string) [][]int {
	// read input as list
	quizs := make([][]int, len(strings.Split(input, "\n")))
	for k, line := range strings.Split(input, "\n") {

		liststr := strings.Fields(line)
		list := make([]int, len(liststr))

		for idx, v := range liststr {
			list[idx], _ = strconv.Atoi(v)
		}
		quizs[k] = make([]int, len(list))
		quizs[k] = list
	}
	return quizs
}
func readAsPyramid(quiz []int) [][]int {
	//fmt.Println(quizs)

	//init
	currPyramid := make([][]int, 1)
	currPyramid[0] = make([]int, len(quiz))
	currPyramid[0] = quiz
	levelCtr := 1
	// build up the pyramid
	for true {
		// init a list to build this level
		// the length = last level len -1
		currList := make([]int, len(currPyramid[levelCtr-1])-1)
		// fill the list
		for idx := 0; idx < len(currPyramid[levelCtr-1])-1; idx++ {
			currList[idx] = currPyramid[levelCtr-1][idx+1] - currPyramid[levelCtr-1][idx]
		}
		// append back
		currPyramid = append(currPyramid, currList)
		// check if the list is convergent to the same number
		// or only one number left
		if convergenceValidator(currList) || len(currList) == 1 {
			break
		} else {
			// continue to next level
			levelCtr++
		}
	}
	//fmt.Println(currPyramid)

	return currPyramid
}

func convergenceValidator(list []int) bool {
	// check if the list is convergent to the same number
	dict := make(map[int]bool)
	for _, v := range list {
		dict[v] = true
	}
	if len(dict) == 1 {
		return true
	}

	return false
}
