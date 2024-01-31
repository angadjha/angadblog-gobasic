package main

import "fmt"

func main() {
	// Decision-making with if statements

	// Example 1: Checking if a number is positive, negative, or zero
	number := 10

	if number > 0 {
		fmt.Println("Number is positive")
	} else if number < 0 {
		fmt.Println("Number is negative")
	} else {
		fmt.Println("Number is zero")
	}

	// Loops including for loops and range loops

	// Example 2: Using a for loop to print numbers from 1 to 5
	fmt.Println("Numbers from 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Example 3: Using a range loop to iterate over elements of a slice
	fruits := []string{"apple", "banana", "orange", "grape", "mango"}
	fmt.Println("Fruits:")
	for index, fruit := range fruits {
		fmt.Printf("%d: %s\n", index+1, fruit)
	}

	// Introduce switch statements for multi-way branching

	// Example 4: Using a switch statement to determine the season based on the month
	month := "January"

	switch month {
	case "January", "February", "December":
		fmt.Println("It's winter")
	case "March", "April", "May":
		fmt.Println("It's spring")
	case "June", "July", "August":
		fmt.Println("It's summer")
	case "September", "October", "November":
		fmt.Println("It's autumn")
	default:
		fmt.Println("Invalid month")
	}

	// Example 5: Using a switch statement to categorize the length of a string
	word := "golang"

	switch len(word) {
	case 1:
		fmt.Println("The word has 1 character")
	case 2, 3:
		fmt.Println("The word has between 2 and 3 characters")
	default:
		fmt.Println("The word has more than 3 characters")
	}
}
