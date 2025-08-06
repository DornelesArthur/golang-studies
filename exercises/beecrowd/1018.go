package main
 
import (
    "fmt"
)
 
func main() {

    var money int
	bankNotes := [7]int{100, 50, 20, 10, 5, 2, 1}

	fmt.Scan(&money)

	fmt.Printf("%d\n", money)
	for i := 0; len(bankNotes) > i; i++ {
		fmt.Printf("%d nota(s) de R$ %d,00\n", money/bankNotes[i], bankNotes[i])
		money = money%bankNotes[i]
	}
	
}
