package main

import "os"

func main() {
	_, err := os.Open("DontExist.txt")

	if err != nil {
		panic(err)
	}
}
