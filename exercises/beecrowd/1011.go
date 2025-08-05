package main
 
import (
    "fmt"
	"math"
)
 
func main() {

	const pi = 3.14159
	var radius int

	fmt.Scan(&radius)

	fmt.Printf("VOLUME = %.3f\n", 4.0/3.0 * pi * math.Pow(float64(radius), 3))

}
