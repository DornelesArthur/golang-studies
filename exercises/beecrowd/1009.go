package main
 
import (
    "fmt"
)
 
func main() {

    var sellerName string
	var salary, salesTotal float64

	fmt.Scan(&sellerName)
	fmt.Scan(&salary)
	fmt.Scan(&salesTotal)

	fmt.Printf("TOTAL = R$ %.2f\n", salary + salesTotal*15/100)
}
