package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	p2()

}

func p1() {
	// fmt.Scanln("%s",&inputPuzzel)
	// check the input value
	// fmt.Printf("Your input is \n%s", inputPuzzel)

	// use While loop process each line by \n
	inputPuzzel := ""
	results := []int{}
	foo := 1
	for foo >= 1 {
		// read stdin until EOF (ctrl + D)
		_, err := fmt.Scanln(&inputPuzzel)
		if err == io.EOF {
			break
		}
		result := ""
		fmt.Printf("The input puzzel is %s\n", inputPuzzel)

		for i := 0; i < len(inputPuzzel); i++ {
			// memory address
			fmt.Print(inputPuzzel[i], ", ")
			// print char
			fmt.Printf("%c \n", inputPuzzel[i])
			if inputPuzzel[i] == 'i' {
				fmt.Printf("6666 gocha")
			}

			// compare the digits and chars
			// if the char is not a digit

			if inputPuzzel[i] > '9' || inputPuzzel[i] < '0' {
				// ignore it
			} else {
				// it's digits, store to result
				result = result + string(inputPuzzel[i])
			}

		}
		// take the first and last digit
		result = string(result[0]) + string(result[len(result)-1])

		num, _ := strconv.Atoi(result)
		results = append(results, num)
		if result != "" {
			fmt.Println("Yes combine digit is ", num)
		}
		result = ""

	}
	sum := 0
	for _, item := range results {
		sum = sum + item
	}
	fmt.Println(sum)

}

func p2() {
	dict := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	// fmt.Println(dict)

	inputFile, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("file err")
		return
	}

	inputString := string(inputFile)
	// fmt.Println(inputString)

	// split by \n
	inputStringArr := strings.Split(inputString, "\n")

	// find the first & last  number for each line
	fisrtidx, lastidx := 9999999, 0
	fisrtNum, lastNum := -1, -1
	ans := 0
	for ctr, str := range inputStringArr {
		for idx, value := range dict {
			// match digits
			if strings.Index(str, strconv.Itoa(idx)) < fisrtidx && strings.Index(str, strconv.Itoa(idx)) != -1 {
				fisrtidx = strings.Index(str, strconv.Itoa(idx))
				fisrtNum = idx
			}

			// match number of string
			if strings.Index(str, value) < fisrtidx && strings.Index(str, value) != -1 {
				fisrtidx = strings.Index(str, value)
				fisrtNum = idx
			}

			// match last digits
			if strings.LastIndex(str, strconv.Itoa(idx)) > lastidx && strings.LastIndex(str, strconv.Itoa(idx)) != -1 {
				lastidx = strings.LastIndex(str, strconv.Itoa(idx))
				lastNum = idx
			}

			// match last number of string
			if strings.LastIndex(str, value) > lastidx && strings.LastIndex(str, value) != -1 {
				lastidx = strings.LastIndex(str, value)
				lastNum = idx
			}
		}
		// got the leftend right end nummber. combine them
		combined := strconv.Itoa(fisrtNum) + strconv.Itoa(lastNum)
		num, _ := strconv.Atoi(combined)
		fmt.Println(ctr, num)
		ans += num
		// init
		fisrtidx, lastidx = 9999999, -1
		fisrtNum, lastNum = -1, -1
	}
	fmt.Println(ans)

	// fmt.Println(inputStringArr)

}
