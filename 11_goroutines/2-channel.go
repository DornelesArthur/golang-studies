package main

import (
	"fmt"
	"time"
)

func main() {
	// Channel buffer
	channel := make(chan int, 100)
	go getIDs(channel)

	// value := <-channel
	// fmt.Println(value)

	for v := range channel {
		fmt.Println("receiving: ", v)
		time.Sleep(time.Second)
	}
}

// Read only channel
// func getIDs(channel <-chan int) {
// "Write" only channel
func getIDs(channel chan<- int) {
	for v := range 100 {
		fmt.Println("sending: ", v)
		channel <- v
	}
	close(channel)
}
