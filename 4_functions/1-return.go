package main

import (
	"fmt"
)

func returnMultipleValues(number1 int, number2 int) (int, int, int, float32) {
	sum := number1 + number2
	sub := number1 - number2
	multi := number1 * number2
	div := float32(number1) / float32(number2)

	return sum, sub, multi, div
}
func main() {
	sum, sub, multi, div := returnMultipleValues(1, 3)
	fmt.Println(sum, sub, multi, div)
}
