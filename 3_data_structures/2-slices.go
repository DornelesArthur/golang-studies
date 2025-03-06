package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Slices x Array
	var array = [5]int{0, 1, 3, 4, 8}
	fmt.Println(array)
	fmt.Println(reflect.TypeOf(array))
	// Error: First Argument to append must be a slice
	// array = append(array,65)

	var slice = []int{0, 1, 3, 4, 8}
	fmt.Println(slice)
	slice = append(slice, 50)
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))

}
