package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) // Creating a channel
	go printMethod(ch)   // Start goroutine
	ch <- 2              // Send a value to the channel

	// Ensure the main goroutine waits for the printMethod goroutine to complete
	time.Sleep(2 * time.Second)
	fmt.Println("Bye from main")
}

func printMethod(ch chan int) {
	// Receive the value sent from main
	val := <-ch

	for i := 0; i < 3; i++ {
		// Perform some operation with the received value
		fmt.Printf("Hello Print Method %v, %v\n", i, val)
		time.Sleep(time.Second)
	}
}
