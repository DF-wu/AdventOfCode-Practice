package main

import "fmt"

func main() {
	// for the puzzel input
	var inputPuzzel string

	// fmt.Scanf("&s", &inputPuzzel)
	// foo := 1
	// for foo>1 {
	//     fmt.Scanln()
	// }

	var a byte = 'a'
	var zero byte = '0'
	fmt.Println("")
	fmt.Println("a:", a)   //UTF-8 : 97(Decimal)
	fmt.Println("0: ", zero) //UTF-8 : 48(Decimal)
	fmt.Println("")
	fmt.Println("Your input is \n%s", &inputPuzzel)

}
