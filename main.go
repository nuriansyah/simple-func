package main

import "fmt"

func main() {
	var inputString1, inputString2 string
	readPrint(inputString1, inputString2)
	appendStringAndPrint()
	printAppendedStringXTimes()
}
func readPrint(inputString1, inputString2 string) {
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

func printAppendedStringXTimes() {
	var inputString1, inputString2 string
	var x int

	fmt.Println("Enter your value")
	fmt.Scan(&x)
	fmt.Print("Enter the first string: ")
	fmt.Scanln(&inputString1)
	fmt.Print("Enter the second string: ")
	fmt.Scanln(&inputString2)
	appendedString := inputString1 + inputString2

	for i := 0; i < x; i++ {
		fmt.Println(appendedString)
	}
}
