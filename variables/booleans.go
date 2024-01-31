package main

import "fmt"

func main() {
	// Declaring and initializing boolean variables
	isSunny := true
	isRainy := false

	// Printing the booleans
	fmt.Println("Is it sunny?", isSunny)
	fmt.Println("Is it rainy?", isRainy)

	// Using boolean values in conditional statements
	if isSunny {
		fmt.Println("Enjoy the sunshine!")
	} else {
		fmt.Println("Hope it clears up soon!")
	}
}
