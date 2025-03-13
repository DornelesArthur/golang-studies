package model

// To be accessible by different packages it must have the first letter capitalized
type Address struct {
	Street string
	Number int
	City   string
}
