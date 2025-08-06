package main
 
import (
    "fmt"
	"math"
)

func commonDivisor(number1 int, number2 int) (cd int){
	for number2 != 0 {
		trash := number1%number2
		number1 = number2
		number2 = trash
	}
	cd = int(math.Abs(float64(number1)))
	return
}

func sum(n1 int, d1 int, n2 int, d2 int){
	d := d1 * d2
	n := (n1*d2) + (n2*d1)
	cd := commonDivisor(n,d)

	fmt.Printf("%d/%d = %d/%d\n", n, d, n/cd, d/cd)
}

func sub(n1 int, d1 int, n2 int, d2 int){
	d := d1 * d2
	n := (n1*d2) - (n2*d1)
	cd := commonDivisor(n,d)

	fmt.Printf("%d/%d = %d/%d\n", n, d, n/cd, d/cd)
}

func mult(n1 int, d1 int, n2 int, d2 int){
	d := d1 * d2
	n := n1 * n2
	cd := commonDivisor(n,d)

	fmt.Printf("%d/%d = %d/%d\n", n, d, n/cd, d/cd)
}

func div(n1 int, d1 int, n2 int, d2 int){
	d := n2 * d1
	n := n1 * d2
	cd := commonDivisor(n,d)

	fmt.Printf("%d/%d = %d/%d\n", n, d, n/cd, d/cd)
}
 
func main() {

    var numberTestCases int
	fmt.Scan(&numberTestCases)

	for i:= 0; i < numberTestCases; i++ {
		var N1, D1, N2, D2 int
		var operator rune

		fmt.Scanf("%d / %d %c %d / %d", &N1, &D1, &operator, &N2, &D2)

		switch operator {
			case '+':
				sum(N1,D1,N2,D2)
			case '-':
				sub(N1,D1,N2,D2)
			case '/':
				div(N1,D1,N2,D2)
			case '*':
				mult(N1,D1,N2,D2)
		}
	}

}
