package main

import "fmt"

// Named function with parameters and a return value
func add(a, b int) int {
	return a + b
}

// Named function with parameters and multiple return values
func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return dividend / divisor, nil
}

func main() {
	// Calling a named function and storing the result
	sum := add(5, 3)
	fmt.Println("Sum:", sum)

	// Calling a named function with multiple return values
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", quotient)
	}

	// Anonymous function assigned to a variable
	multiply := func(x, y int) int {
		return x * y
	}

	// Calling an anonymous function through the variable
	product := multiply(4, 6)
	fmt.Println("Product:", product)
}
