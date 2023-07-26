package main

import "fmt"

func main() {
	readPrint()
	appendStringAndPrint()
}
func readPrint() {
	var inputString1, inputString2 string
	fmt.Print("Enter the first string: ")
	fmt.Scanln(&inputString1)
	fmt.Print("Enter the second string: ")
	fmt.Scanln(&inputString2)
	fmt.Println(inputString1, inputString2)
}

func appendStringAndPrint() {
	var inputString1, inputString2 string

	fmt.Print("Enter the first string: ")
	fmt.Scanln(&inputString1)
	fmt.Print("Enter the second string: ")
	fmt.Scanln(&inputString2)

	appendedString := inputString1 + inputString2

	fmt.Println("Appended string:", appendedString)
}
