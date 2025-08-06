package main
 
import (
    "fmt"
)
 
func main() {

	var ageInDays int

	fmt.Scan(&ageInDays)

	fmt.Printf("%d ano(s)\n", ageInDays/365)
	fmt.Printf("%d mes(es)\n", (ageInDays%365)/30)
	fmt.Printf("%d dia(s)\n",(ageInDays%365)%30)

}
