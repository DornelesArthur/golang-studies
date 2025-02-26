package main

import "fmt"

func main() {
	salary := 1201.00
	contract_type := "PJ"

	if salary < 1000 || contract_type == "CLT" {
		salary -= (salary * 0.08)
	} else if salary <= 1200 && contract_type == "PJ" {
		salary -= (salary * 0.05)
	} else if salary <= 1200.00 {
		salary -= (salary * 0.10)
	} else {
		salary -= (salary * 0.15)
	}
	fmt.Println("Salary: ", salary)
}
