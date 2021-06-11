package main

import "fmt"

func ChInAction() {
	c := make(chan bool)
	go waitAndsay(c, "World")
	fmt.Println("Hello")

	//We send a signal to c inorder for waitAndSay to continue
	c <- true

	//We wait to receive another signal on c before we exit

	<-c

}

func waitAndsay(c chan bool, s string) {
	if b := <-c; b {
		fmt.Println(s)
	}
	// <-c will receive signal on channel to exit
	c <- true
}
