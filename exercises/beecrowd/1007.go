package main
 
import (
    "fmt"
)
 
func main() {

    var A, B, C, D int
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	fmt.Scan(&D)

	fmt.Printf("DIFERENCA = %d\n", A*B-C*D)

}
