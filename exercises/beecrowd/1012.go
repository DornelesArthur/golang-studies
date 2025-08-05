package main
 
import (
    "fmt"
	"math"
)
 
func main() {

	const pi = 3.14159
    var A, B, C float64

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)

	fmt.Printf("TRIANGULO: %.3f\n", A*C/2)
	fmt.Printf("CIRCULO: %.3f\n", pi * math.Pow(float64(C), 2))
	fmt.Printf("TRAPEZIO: %.3f\n", (A+B) * C/2)
	fmt.Printf("QUADRADO: %.3f\n", B*B)
	fmt.Printf("RETANGULO: %.3f\n", A*B)

}
