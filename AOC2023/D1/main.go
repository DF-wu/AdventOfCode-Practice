package main

import (
	"fmt"
	"io"
	"strconv"
)

func main() {

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
		result = string(result[0]) + string(result[len(result) - 1])

		num, _ := strconv.Atoi(result)
		results = append(results, num)
		if result != "" {
			fmt.Println("Yes combine digit is ", num)
		}
		result = ""

	}
	sum := 0
	for _ , item := range results{
		sum = sum + item
	}
	fmt.Println(sum)

}
