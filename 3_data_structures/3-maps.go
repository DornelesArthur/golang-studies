package main

import (
	"fmt"
)

func main() {
	// cities_population := map[string]int{"São Paulo": 10000000, "New York": 8000000, "Vancouver": 2642825}
	cities_population := make(map[string]int)
	cities_population["São Paulo"] = 10000000
	cities_population["New York"] = 8000000
	cities_population["Vancouver"] = 2642825
	fmt.Println(cities_population)

	if value, key_exists := cities_population["Paris"]; key_exists {
		fmt.Println(value)
	} else {
		fmt.Println("Key dont exist")
	}

	delete(cities_population, "New York")
	delete(cities_population, "Paris")

	for key, value := range cities_population {
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", value)
	}
}
