package main

import "fmt"

func main() {
	fmt.Println("Hello main")
	message := make(chan string)
	go func() { message <- "Hello Chan" }()
	msg := <-message
	fmt.Println("message: ", msg)
}
