package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wait_group sync.WaitGroup
	// goroutine is similar to threads (goroutines typically use 2kb of ram, instead of 2mb like threads)
	// Go Scheduler decides (MN Scheduling) which goroutines will be in the same thread
	for i := range 10 {
		wait_group.Add(2)
		// Just use go to call func async
		go showMessage("Hello World! "+strconv.Itoa(i), &wait_group)
		go showMessage2("Message2: "+strconv.Itoa(i), &wait_group)
	}
	wait_group.Wait()
}

func showMessage(message string, wait_group *sync.WaitGroup) {
	fmt.Println(message)
	wait_group.Done()
}

func showMessage2(message string, wait_group *sync.WaitGroup) {
	fmt.Println(message)
	wait_group.Done()
}
