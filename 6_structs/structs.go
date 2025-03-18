package main

import (
	"fmt"
	model_alias "golang_studies/6_structs/model"
	"time"
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
		BirthDate: time.Date(1995, 03, 18, 0, 0, 0, 0, time.Local),
	}
	fmt.Println(address)
	fmt.Println(person)
	fmt.Println(model_alias.GetAge(person))
	fmt.Println(person.GetAge())
	person.SetHeigth()
	fmt.Println(person.Height)
}
