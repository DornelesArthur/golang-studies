package main
 
import (
    "fmt"
	"math"
)
 
func main() {

    var X1, Y1, X2, Y2, distance float64

	fmt.Scan(&X1)
	fmt.Scan(&Y1)
	fmt.Scan(&X2)
	fmt.Scan(&Y2)

	distance = math.Sqrt(math.Pow(X2-X1, 2.0) + math.Pow(Y2-Y1, 2.0))
	
	fmt.Printf("%.4f\n", distance)
}
