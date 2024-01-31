package main

import "fmt"

func main() {
	// Introducing arrays and slices

	// Arrays: Fixed-size collection of elements of the same type
	var array [5]int // Declare an array of size 5 to store integers
	array[0] = 1     // Assigning values to array elements
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5

	// Slices: Dynamic-size view into the elements of an array
	// Syntax: slice[startIndex:endIndex]
	slice := array[1:4] // Create a slice from index 1 to 3 (endIndex-1) of the array
	// The slice references the elements [2 3 4] from the array

	// Covering basic operations on slices
	fmt.Println("Original Array:", array)
	fmt.Println("Slice:", slice)

	// Appending elements to a slice
	slice = append(slice, 6, 7)
	fmt.Println("After Appending:", slice)

	// Accessing elements of a slice
	fmt.Println("First Element:", slice[0])
	fmt.Println("Last Element:", slice[len(slice)-1])

	// Introducing maps

	// Maps: Unordered collection of key-value pairs
	// Syntax: map[keyType]valueType
	studentAges := map[string]int{
		"John":  25,
		"Alice": 30,
		"Bob":   28,
	}

	// Demonstrating usage of maps
	fmt.Println("Student Ages:", studentAges)

	// Accessing elements of a map
	fmt.Println("Age of John:", studentAges["John"])

	// Adding a new key-value pair to the map
	studentAges["Emily"] = 26
	fmt.Println("Student Ages (after adding Emily):", studentAges)

	// Deleting a key-value pair from the map
	delete(studentAges, "Bob")
	fmt.Println("Student Ages (after deleting Bob):", studentAges)
}
