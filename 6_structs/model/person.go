package model

import (
	"time"
)

type Person struct {
	Name      string
	Address   Address
	BirthDate time.Time
	Height    int
}

// "method"
func (p Person) GetAge() (age int) {
	age = time.Now().Year() - p.BirthDate.Year()
	if time.Now().Month() < p.BirthDate.Month() {
		age -= 1
	} else if time.Now().Month() == p.BirthDate.Month() && time.Now().Day() < p.BirthDate.Day() {
		age -= 1
	}
	return
}

// normal func
func GetAge(p Person) (age int) {
	age = time.Now().Year() - p.BirthDate.Year()
	if time.Now().Month() < p.BirthDate.Month() {
		age -= 1
	} else if time.Now().Month() == p.BirthDate.Month() && time.Now().Day() < p.BirthDate.Day() {
		age -= 1
	}
	return
}

// Must be a pointer
func (p *Person) SetHeigth() {
	p.Height = 175
}
