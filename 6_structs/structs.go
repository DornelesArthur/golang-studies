package main

import (
	"fmt"
	"golang_studies/6_structs/address"
)

func main() {
	fmt.Println("Init")

	address := address.Address{
		Street: "2A Macquarie Street",
		Number: 12512,
		City:   "Atlantis",
	}
	fmt.Println(address)
}
