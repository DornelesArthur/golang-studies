package main
 
import (
    "fmt"
)
 
func main() {

    var tripTime, averageSpeed int

	fmt.Scan(&tripTime)
	fmt.Scan(&averageSpeed)

	fmt.Printf("%.3f\n", float64(tripTime*averageSpeed)/12)
}
