package main

import "fmt"

func main() {
	string_array := [5]string{"test", "test", "test", "test", "test"}

	fmt.Println(string_array)
	fmt.Println(len(string_array))

	int_array := [5]int{1, 3, 2, 5, 10}

	fmt.Println(int_array)
	fmt.Println(len(int_array))

	var float_array [3]float64
	pi := 3.1415

	for i := 0; i < len(float_array); i++ {
		float_array[i] = pi + float64(i)
	}
	fmt.Println(float_array)
	fmt.Println(len(float_array))

	bool_array := [3]bool{true, true, false}
	for index, value := range bool_array {
		fmt.Println(value)
		fmt.Println(index)
	}
	fmt.Println(len(bool_array))

	// append
	list := []int{4, 5, 6, 8, 10}

	fmt.Println(list)
	list = append(list, 16)
	fmt.Println(list)

	// make
	new_list := make([]int, 1)
	new_list = append(new_list, 23)
	fmt.Println(new_list)

	new_list2 := list[:3]
	new_list3 := list[3:]
	last_number := list[len(list)-1:]
	fmt.Println(new_list2)
	fmt.Println(new_list3)
	fmt.Println(last_number)
}
