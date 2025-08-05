package main
 
import (
    "fmt"
)
 
func main() {
	var A, B, C, MEAN float64
    
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	MEAN = (A*2 + B*3 + C*5)/(10)
	fmt.Printf("MEDIA = %.1f\n", MEAN)

}
