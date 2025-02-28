package main

import (
	"fmt"
)

func bubbleSort(books_ids []string, books_total int) {
	for i := 0; i < books_total; i++ {
		for j := 0; j < books_total-1; j++ {
			if books_ids[j] > books_ids[j+1] {
				x := books_ids[j]
				books_ids[j] = books_ids[j+1]
				books_ids[j+1] = x
			}
		}
	}
}

func main() {
	var books_total int
	var x string

	for {
		_, err := fmt.Scanln(&books_total)
		books_ids := make([]string, books_total)
		if err != nil {
			return
		}
		for i := 0; i < books_total; i++ {
			fmt.Scan(&x)
			books_ids[i] = x
		}
		bubbleSort(books_ids, books_total)

		for i := 0; i < books_total; i++ {
			fmt.Println(books_ids[i])
		}
	}
}
