package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
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

	var mutex sync.Mutex
	number := 0

	for x := 0; x < 10; x++ {
		wait_group.Add(1)
		go func() {
			mutex.Lock()
			number = x
			time.Sleep(time.Second * 3)
			// Without Mutex any gorountine can change the value of number at any time and will enter the if and exec panic()
			if x != number {
				panic("Number is the same as X")
			}
			mutex.Unlock()
			wait_group.Done()
		}()
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
