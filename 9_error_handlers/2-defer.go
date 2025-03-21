package main

import (
	"fmt"
	"os"
)

func ShowText() {
	fmt.Println("Finish!!")
}
func main() {
	file, err := os.Create("./settings.txt")
	defer file.Close()
	defer ShowText()

	if err != nil {
		panic(err)
	}
	fmt.Println("Writing")
	file.Write([]byte("Test"))

}
