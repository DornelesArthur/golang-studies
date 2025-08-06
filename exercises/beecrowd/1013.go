package main

import (
    "fmt"
	"math"
)
 
func main() {

	var A, B, C int 
	var highest float64
    fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)

	highest = (float64(A+B)+math.Abs(float64(A-B)))/2
	highest = (highest+float64(C)+math.Abs(highest-float64(C)))/2

	fmt.Printf("%.0f eh o maior\n", highest)
}