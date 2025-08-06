package main
 
import (
    "fmt"
)
 
func main() {

    var seconds int

	fmt.Scan(&seconds)

	fmt.Printf("%d:", seconds/(60*60))
	seconds = seconds%(60*60)
	fmt.Printf("%d:", seconds/60)
	seconds = seconds%60
	fmt.Printf("%d\n", seconds)
}
