package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- "Form1"
		}
	}()
	go func() {
		for {
			ch2 <- "Form2"
		}
	}()

	for {
		select {
		case msg1 := <-ch1:
			fmt.Printf(msg1)
		case msg2 := <-ch2:
			fmt.Printf(msg2)
		default:
			fmt.Printf("no message recived")
		}
	}
}
