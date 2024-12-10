package main

import (
	"fmt"
	"math"
)

func main() {
	var A float64

	fmt.Scan(&A)
	fmt.Printf("A=%0.4f\n", math.Pow(A, 2)*3.14159)
}
