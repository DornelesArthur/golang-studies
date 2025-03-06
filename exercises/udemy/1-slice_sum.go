package main

import (
	"fmt"
)

func main() {
	my_slice := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}
	sum_number_below_6 := 0
	sum_number_above_5 := 0

	for _, value := range my_slice {
		fmt.Println(value)
		if value < 6 {
			sum_number_below_6 += value
		} else {
			fmt.Println(value)
			sum_number_above_5 += value
		}
	}
	fmt.Println(sum_number_below_6)
	fmt.Println(sum_number_above_5)

}
