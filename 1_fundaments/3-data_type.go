package main

import "fmt"

func main() {
	// := declare variable and set value  
	text := "Text"
	// Variables data types cannot change 
	// text = 10

	text = "Update Text"
	fmt.Println(text)

	// Primitive Types
	var var_integer int
	var var_string string
	var var_float32 float32
	var var_bool bool

	// Primitive Types default values
	fmt.Println(var_integer)
	fmt.Println(var_string)
	fmt.Println(var_float32)
	fmt.Println(var_bool)
}