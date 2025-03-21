package main

import "fmt"

type test struct {
	name string
	id   int
}

type customConstraint interface {
	int | string | bool | test
}

func main() {
	sliceInt := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(reverse(sliceInt))

	sliceString := []string{"a", "b", "asda", "hse", "z"}
	fmt.Println(reverse(sliceString))

	sliceBool := []bool{false, true, true, true, false, true}
	fmt.Println(reverse(sliceBool))

	sliceTest := []test{{name: "dfassd", id: 1}, {name: "23521", id: 3}, {name: "Xsafa123", id: 123}}
	fmt.Println(reverse(sliceTest))
}

func reverse[T customConstraint](slice []T) (newSlice []T) {
	newSlice = make([]T, len(slice))
	for i, j := 0, len(slice)-1; i < len(slice); i, j = i+1, j-1 {
		newSlice[j] = slice[i]
	}
	return
}
