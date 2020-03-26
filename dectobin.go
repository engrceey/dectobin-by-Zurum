package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Variable declaration
	var userInput int64

	// Message to user
	fmt.Println("Enter a decimal number below >> ")

	// Read in user input to memory
	fmt.Scanln(&userInput)

	// Convert Decimal input to Binary
	answer := strconv.FormatInt(userInput, 2)

	// Print Answer
	fmt.Printf("Input Number = %d \nConverted Number = %s \n\n", userInput, answer)

}
