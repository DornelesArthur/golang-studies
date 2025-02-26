package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	/* i := 10
	for ; i >= 1; i-- {
		fmt.Println(i)
	} */

	text := "Text"
	for i := 0; len(text) > i; i++ {
		if string(text[i]) == "x" {
			break
		}
		fmt.Println(string(text[i]))
	}

	word := "pneumonoultramicroscopicsilicovolcanoconiosis"
	word_length := len(word)
	i := 0
	for word_length > i {
		if string(word[i]) == "t" {
			break
		}
		fmt.Println(string(word[i]))
		i++
	}

}
