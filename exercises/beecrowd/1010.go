package main
 
import (
    "fmt"
)
 
func main() {

    var productID1, productID2, amoutProduct1, amoutProduct2 int
	var priceProduct1, priceProduct2 float64

	fmt.Scan(&productID1)
	fmt.Scan(&amoutProduct1)
	fmt.Scan(&priceProduct1)
	fmt.Scan(&productID2)
	fmt.Scan(&amoutProduct2)
	fmt.Scan(&priceProduct2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", float64(amoutProduct1)*priceProduct1 + float64(amoutProduct2)*priceProduct2)
}
