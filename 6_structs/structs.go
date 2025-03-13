package main

import (
	"fmt"
	address_alias "golang_studies/6_structs/address"
)

func main() {
	fmt.Println("Init")

	address := address_alias.Address{
		Street: "2A Macquarie Street",
		Number: 12512,
		City:   "Atlantis",
	}
	fmt.Println(address)
}
