package main

import (
	"fmt"
	model_alias "golang_studies/6_structs/model"
)

func main() {
	fmt.Println("Init")

	address := model_alias.Address{
		Street: "2A Macquarie Street",
		Number: 12512,
		City:   "Atlantis",
	}

	person := model_alias.Person{
		Name: "Angrboda",
		Address: model_alias.Address{
			Street: "Not 2A Macquarie Street",
			Number: 21521,
			City:   "Not Atlantis",
		},
	}
	fmt.Println(address)
	fmt.Println(person)
}
