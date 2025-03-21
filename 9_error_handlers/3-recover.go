package main

import (
	"fmt"
	"os"
)

func ReadFile() {
	// Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution. (https://go.dev/blog/defer-panic-and-recover)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered")
		}
	}()

	_, err := os.Open("DontExist.txt")
	if err != nil {
		panic(err)
	}
}
func main() {
	ReadFile()

	fmt.Println("Finish")
}
