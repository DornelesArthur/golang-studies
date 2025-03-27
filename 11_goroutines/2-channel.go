package main

import "fmt"

func main() {
	channel := make(chan int)
	go getIDs(channel)

	// value := <-channel
	// fmt.Println(value)

	for v := range channel {
		fmt.Println(v)
	}
}

func getIDs(channel chan int) {
	for v := range 100 {
		channel <- v
	}
	close(channel)
}
