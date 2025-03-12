package main

import (
	"fmt"
)

func testPointers(pointer *int, pointer2 *int) {
	fmt.Println("TEST Pointers")
	*pointer = 1
	*pointer2 = 2
	fmt.Println(pointer, pointer2)
}

func main() {
	value1 := 15
	value2 := 15
	fmt.Println(value1, value2)
	fmt.Println(&value1, &value2)

	value1_pointer := &value1
	*value1_pointer = 163

	fmt.Println(value1, value2)
	fmt.Println(&value1, &value2, value1_pointer, &value1_pointer)
	testPointers(&value1, &value2)

	fmt.Println("After TEST Pointers")
	fmt.Println(value1, value2)
	fmt.Println(&value1, &value2)
}
