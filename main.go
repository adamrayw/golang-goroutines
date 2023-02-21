package main

import (
	"fmt"
	"time"
)

// It prints the numbers 0 to 4, with a half second delay between each number
func print() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(500 * time.Millisecond)
	}
}

// PrintLetters prints the letters a through d, pausing for half a second between each letter
func printLetters() {
	for i := 'a'; i < 'e'; i++ {
		fmt.Printf("%c ", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// Calling the function print() in a new goroutine.
	go print()
	go printLetters()
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
}
