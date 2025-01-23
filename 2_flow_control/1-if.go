package main

import "fmt"

func main() {
	salary := 1201.00

	if salary < 1000 {
		salary -= (salary * 0.08)
	} else if salary <= 1200 {
		salary -= (salary * 0.1)
	} else {
		salary -= (salary * 0.15)
	}
	fmt.Println("Salary: ", salary)
}
