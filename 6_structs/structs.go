package main

import (
	"fmt"
)

type address struct {
	street string
	number int
	city   string
}

func main() {
	fmt.Println("Init")

	address := address{
		street: "2A Macquarie Street",
		number: 12512,
		city:   "Atlantis",
	}
	fmt.Println(address)
}
