package main

import (
	"fmt"
	"reflect"
)

func main() {
	const tax = 10
	tax = 5

	fmt.Println(tax)
	fmt.Println(reflect.TypeOf(tax))
}