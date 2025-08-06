package main
 
import (
    "fmt"
)
 
func main() {

    var km int
	var spentFuel float64

	fmt.Scan(&km)
	fmt.Scan(&spentFuel)

	fmt.Printf("%.3f km/l\n", float64(km)/spentFuel)
}
