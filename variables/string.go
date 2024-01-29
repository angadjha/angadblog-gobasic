package main

import "fmt"

const s string = "constant"

func main() {
	var a = "first"
	fmt.Println(a)

	var b, c, f int = 1, 2, 5
	fmt.Println(b, c, f)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	g := "apple"
	fmt.Println(g)

	//Call String method
	callString()
}

func callString() {
	var welcome string = "Hello Angad"
	fmt.Println("String length- ", len(welcome))
	fmt.Println("Print First Character- ", string(welcome[0]))
	fmt.Println("Print last Character- ", string(welcome[len(welcome)-1]))

	// Accessing individual characters (runes) in the string
	for i, char := range welcome {
		fmt.Println("print each character- ", i, char)
	}

	// Concatenating strings
	text := " Welcome to Go Programming!"
	concatenated := welcome + text
	fmt.Println("Concatenated string:", concatenated)

	// slicing string
	str := welcome[0:5]
	fmt.Println("Sliced String 0-4 ", str)

	// checking string contains Angad
	contains := containsString(welcome, "Angad")
	if contains {
		fmt.Println("Value Contains")
	} else {
		fmt.Println("value not contains")
	}
}

func containsString(s, substr string) bool {
	return (len(s) > 0 && len(substr) > 0 && len(s) >= len(substr) &&
		(s == substr || (len(s) > len(substr) && containsString(s[1:], substr))))
}
