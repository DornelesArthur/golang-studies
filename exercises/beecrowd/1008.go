package main
 
import (
    "fmt"
)
 
func main() {

	var workerID, workedHours int
	var salaryPerHour float64

	fmt.Scan(&workerID)
	fmt.Scan(&workedHours)
	fmt.Scan(&salaryPerHour)

	fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n", workerID, float64(workedHours)*salaryPerHour);

}
