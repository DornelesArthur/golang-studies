package main
 
import (
    "fmt"
	"math"
)
 
func main() {

	var moneyInt int
    var money float64

	fmt.Scan(&money)
	moneyNotes, money := math.Modf(money)
	moneyInt = int(moneyNotes)
	money = math.Round(money*100) / 100

	bankNotes := [6]int{100, 50, 20, 10, 5, 2}
	bankCoins := [5]float64{0.5, 0.25, 0.1, 0.05, 0.01}

	fmt.Println("NOTAS:")
	for i := 0; len(bankNotes) > i; i++ {
		fmt.Printf("%d nota(s) de R$ %d.00\n", moneyInt/bankNotes[i], bankNotes[i])
		moneyInt = moneyInt%bankNotes[i]
	}

	fmt.Println("MOEDAS:")
	fmt.Printf("%d moeda(s) de R$ 1.00\n", moneyInt)
	for i := 0; len(bankCoins) > i; i++ {
		quantity := int(money/bankCoins[i])
		fmt.Printf("%d moeda(s) de R$ %.2f\n", quantity, bankCoins[i])
		money = money - float64(quantity) * bankCoins[i]
		money = math.Round(money*100) / 100
	}
}
