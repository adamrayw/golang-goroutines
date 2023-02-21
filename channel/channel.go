package main

import "fmt"

func SendMessage(messages chan string) {
	messages <- "Hello"
	messages <- "World"
	close(messages)
}

func main() {
	// Calling the function print() in a new goroutine.
	messages := make(chan string)

	go SendMessage(messages)

	for message := range messages {
		fmt.Println(message)
	}
}
