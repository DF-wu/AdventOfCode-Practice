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

}

func p1(puzzle string) {
	// puzzle to 2d array
	puzzles := strings.Split(puzzle, "\n")
	for k, _ := range puzzles {
		puzzles[k] = strings.Trim(puzzles[k], "\n")
	}
	// engine := make([][]string, len(puzzles))

	ans := 0

	for row, line := range puzzles {
		numflag := false
		leftidx, rightidx := 0, 0
		for col, char := range line {
			numflagstatus := numflag
			switch char {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				// number, find left index and right index
				numflag = true
			default:
				numflag = false
			}
			if numflagstatus != numflag {
				// pivot pass over left index or right index
				if numflag {
					// left index
					leftidx = col
				} else {
					// right index
					rightidx = col
					// verify engine part
					if enginePartValidator(puzzles, row, leftidx, rightidx) {
						// find a valid engine part
						// parse the number and add to sum
						numberchars := puzzles[row][leftidx:rightidx]
						// fmt.Println(numberchars)
						inum, _ := strconv.Atoi(numberchars)
						ans += inum
					} else {
						// str := puzzl es[row][leftidx:rightidx]
						// fmt.Println(str)
					}
				}
			}
		}
	}
	fmt.Println(ans)

}

func enginePartValidator(puzzles []string, row int, leftidx int, rightidx int) bool {
	/* if select 633, a is the elements to check
	467..114..
	...*.aaaaa
	..35.a633a
	.....a#aaa
	617*......
	.....+.58.
	*/

	for i := row - 1; i <= row+1; i++ {
		if i < 0 || i >= len(puzzles) {
			// out of bound
			continue
		}
		for j := leftidx - 1; j <= rightidx+1; j++ {
			if j < 0 || j >= len(puzzles[i]) {
				// out of bound
				continue
			}
			if (puzzles[i][j] >= '0' && puzzles[i][j] <= '9') || puzzles[i][j] == '.' {
				// is a number, pass it
				// . means nothing
				// ignore ascii code control codes
			} else if puzzles[i][j] < 33 {
				fmt.Printf("!!! %v\n", puzzles[i][j])

			} else {
				// got symbol = is a part of engine
				// fmt.Printf("the rune is %v | ", string(puzzles[i][j]))
				return true
			}

		}

	}
	return false
}
