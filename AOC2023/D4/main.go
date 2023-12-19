package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	id      int
	winNums []int
	nums    []int
}

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

	p2(input)
}

func p2(puzzle string) {
	// deal with deserialization first
	//   in Windows -> CR + LF
	puzzles := strings.Split(puzzle, "\r\n")
	cards := []Card{}
	for k, v := range puzzles {
		//  seperate the numbers
		thenums := strings.Split(strings.Split(v, ":")[1], "|")

		// parse the numbers form string to int
		wnumsStr := strings.Fields(thenums[0])
		wnums := make([]int, len(wnumsStr))
		for i := 0; i < len(wnumsStr); i++ {
			tmp, _ := strconv.Atoi(wnumsStr[i])
			wnums[i] = tmp
			//  below is wrong, because the wnums is slice but not the map
			// wnums[i], _ = strconv.Atoi(strings.Fields(thenums[0])[i])
		}

		elsenumsStr := strings.Fields(thenums[1])
		elsenums := make([]int, len(elsenumsStr))
		for k, v := range elsenumsStr {
			tmp, _ := strconv.Atoi(v)
			elsenums[k] = tmp
		}

		c := Card{
			id:      k,
			winNums: wnums,
			nums:    elsenums,
		}
		cards = append(cards, c)

	}

	// solve
	cardsOwn := make([]int, len(cards))
	// init
	for k, _ := range cardsOwn {
		cardsOwn[k] = 1
	}

	for id, v := range cards {
		// init dict for each game
		dict := make(map[int]int)

		// init map
		for _, num := range v.winNums {
			dict[num] = 1
		}
		// match
		for _, num := range v.nums {
			_, isExist := dict[num]
			if isExist {
				dict[num]++
			} else {
				// pass
			}

		}

		ctr := 0
		for _, num := range dict {
			if num == 2 {
				// got matched
				ctr++
			}
		}
		fmt.Println("This round mathed num: ", ctr)

		tmp := cardsOwn[id]

			// increase cards for each match
			for i := 0; i < ctr; i++ {
				cardsOwn[i+id+1] = cardsOwn[i+id+1] + tmp
			}
		
		

		/*
		tmp := cardsOwn[id]
		for rs := 0; rs < tmp; rs++ {
			// increase cards for each match
			for i := 0; i < ctr; i++ {
				cardsOwn[i+id+1] = cardsOwn[i+id+1] + 1
			}
		}
		*/

	}
	fmt.Println(cardsOwn)
	ans := 0
	for _, v := range cardsOwn {
		ans += v
	}
	println(ans)
}

func p1(puzzle string) {
	// deal with deserialization first
	//   in Windows -> CR + LF
	puzzles := strings.Split(puzzle, "\r\n")
	cards := []Card{}
	for k, v := range puzzles {
		//  seperate the numbers
		thenums := strings.Split(strings.Split(v, ":")[1], "|")

		// parse the numbers form string to int
		wnumsStr := strings.Fields(thenums[0])
		wnums := make([]int, len(wnumsStr))
		for i := 0; i < len(wnumsStr); i++ {
			tmp, _ := strconv.Atoi(wnumsStr[i])
			wnums[i] = tmp
			//  below is wrong, because the wnums is slice but not the map
			// wnums[i], _ = strconv.Atoi(strings.Fields(thenums[0])[i])
		}

		elsenumsStr := strings.Fields(thenums[1])
		elsenums := make([]int, len(elsenumsStr))
		for k, v := range elsenumsStr {
			tmp, _ := strconv.Atoi(v)
			elsenums[k] = tmp
		}

		c := Card{
			id:      k,
			winNums: wnums,
			nums:    elsenums,
		}
		cards = append(cards, c)
		fmt.Println(c)

	}

	// solve

	ans := 0
	for _, v := range cards {
		// init dict for each game
		dict := make(map[int]int)

		// init map
		for _, num := range v.winNums {
			dict[num] = 1
		}
		// match
		for _, num := range v.nums {
			_, isExist := dict[num]
			if isExist {
				dict[num]++
			} else {
				// pass
			}

		}
		fmt.Println(dict)
		ctr := 0
		for _, num := range dict {
			if num == 2 {
				// got matched
				ctr++
			}
		}
		fmt.Println("This round ", powInt(2, ctr-1))
		ans += powInt(2, ctr-1)

	}
	fmt.Println(ans)
}

// Golang math.pow() return float64!!!
func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
