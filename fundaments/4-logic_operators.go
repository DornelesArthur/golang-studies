package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("SUM")
	num1 := 10
	num2 := 20
	result := num1 + num2
	fmt.Println(result)

	fmt.Println("SUB")
	result = num1 - num2
	fmt.Println(result)

	fmt.Println("MULTI")
	result = num1 * num2
	fmt.Println(result)

	fmt.Println("DIV INT")
	result = num1 / num2
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
	result = num1 % num2
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))

	fmt.Println("DIV FLOAT")
	result2 := 10.0 / 20.0
	fmt.Println(result2)
	fmt.Println(reflect.TypeOf(result2))

	fmt.Println("CONCAT STRINGS")
	string1 := "He"
	string2 := "llo"
	fmt.Println(string1 + string2)
}