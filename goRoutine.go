package main

import (
	"fmt"
	"time"
)

func runGoRoutine() {
	//word variable changed in calling runGoRoutine function
	//Go routine works from same memory space as of the caller
	word := "Hello"
	//Anonymous function
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println(word)
	}()
	fmt.Println("Hello")
	word = "World"
	time.Sleep(3 * time.Second)

}
