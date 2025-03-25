package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// goroutine is similar to threads (goroutines typically use 2kb of ram, instead of 2mb like threads)
	// Go Scheduler decides (MN Scheduling) which goroutines will be in the same thread
	for i := 0; i < 10; i++ {
		// Just use go to call func async
		go showMessage("Hello World! " + strconv.Itoa(i))
	}

	time.Sleep(time.Duration(time.Hour.Seconds() * float64(100)))
}

func showMessage(message string) {
	fmt.Println(message)
}
