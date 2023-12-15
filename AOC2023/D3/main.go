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
	// p1(input)
	p2(input)
}


func p2(puzzle string){
	
	puzzles := strings.Split(puzzle, "\r\n")

	// engine := make([][]string, len(puzzles))

	adjTable := map[int][]int{}

	for row, line := range puzzles {
		numflag := false
		leftidx, rightidx := 0, 0
		for col, char := range line {
			numflagstatus := numflag
			switch char {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				// number, find left index and right index
				numflag = true
			default:
				numflag = false
			}
			toValidate := false
			// status handle
			// a. start to read number:  including end of line
			// b. reading number excluding last character
			// c. end reading number.   including end of line
			// d. read nothing
			if numflagstatus == false && numflag == true {
				// a
				leftidx = col
				// init. at least one digit
				rightidx = col
				if(col == len(line)-1) { toValidate =true }
			} else if numflagstatus == true && numflag == true {
				// b
				rightidx++
				if(col == len(line)-1) { toValidate =true }
			} else if numflagstatus == true && numflag == false {
				// c
				toValidate = true
			} else {
				//d
				
			}
			if( toValidate ){
				engineGearLocator(puzzles, row, leftidx, rightidx, adjTable) 
			}
			

		}
	}

	ans := 0
	for _, v := range adjTable {
		if(len(v) == 2){
			ans += v[0] * v[1]
		}
	}
	fmt.Println(ans)
}


func engineGearLocator(puzzles []string, row int, leftidx int, rightidx int, adjTable map[int][]int){
	//  use adjecency table store number for a "*"
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
				// ignore ascii code control codes
			} else if puzzles[i][j] == '*'{
				// got symbol -> store to adjecency list
				// cal the position of the '*' and append the number to the list
				numberchars := puzzles[row][leftidx : rightidx+1]
				val, _ := strconv.Atoi(numberchars)
				position := i*len(puzzles[i])+j
				adjTable[position] = append(adjTable[position], val)
				
			}

		}

	}
	
}

func p1(puzzle string) {

	puzzles := strings.Split(puzzle, "\r\n")

	// engine := make([][]string, len(puzzles))

	ans := 0

	for row, line := range puzzles {
		numflag := false
		leftidx, rightidx := 0, 0
		for col, char := range line {
			numflagstatus := numflag
			switch char {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				// number, find left index and right index
				numflag = true
			default:
				numflag = false
			}
			toValidate := false
			// status handle
			// a. start to read number:  including end of line
			// b. reading number excluding last character
			// c. end reading number.   including end of line
			// d. read nothing
			if numflagstatus == false && numflag == true {
				// a
				leftidx = col
				// init. at least one digit
				rightidx = col
				if(col == len(line)-1) { toValidate =true }
			} else if numflagstatus == true && numflag == true {
				// b
				rightidx++
				if(col == len(line)-1) { toValidate =true }
			} else if numflagstatus == true && numflag == false {
				// c
				toValidate = true
			} else {
				//d
				
			}
			if( toValidate ){
				if enginePartValidator(puzzles, row, leftidx, rightidx) {
					// find a valid engine part
					// parse the number and add to sum
					numberchars := puzzles[row][leftidx : rightidx+1]
					// fmt.Println(numberchars)
					inum, _ := strconv.Atoi(numberchars)
					ans += inum
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
				// ignore ascii code control codes
			} else {
				// got symbol = is a part of engine
				// fmt.Printf("the rune is %v | ", string(puzzles[i][j]))
				return true
			}

		}

	}
	return false
}

