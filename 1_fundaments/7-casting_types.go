package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println("Convert int")
	var number8 int8 = 120
	var numberInt int
	numberInt = int(number8)

	fmt.Println(numberInt)
	fmt.Println(reflect.TypeOf(numberInt))
	// fmt.Println('--------------')
	fmt.Println(number8)
	fmt.Println(reflect.TypeOf(number8))

	fmt.Println("Convert float to int")
	var number int
	var numberFloat float32 = 128.99999
	number = int(numberFloat)

	fmt.Println(numberFloat)
	fmt.Println(reflect.TypeOf(numberFloat))
	// fmt.Println('--------------')
	fmt.Println(number)
	fmt.Println(reflect.TypeOf(number))

	fmt.Println("Convert strings")
	boolean, err := strconv.ParseBool("true")
	fmt.Println(boolean)
	fmt.Println(err)

	boolean, _ = strconv.ParseBool("true")
	fmt.Println(boolean)

	boolean, err = strconv.ParseBool("something else")
	fmt.Println(boolean)
	fmt.Println(err)
}