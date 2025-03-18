package model

type Vehicle struct {
	Year         int
	Brand        string
	Model        string
	LicensePlate string
}

type Motorcycle struct {
	Vehicle
	CC int
}

type Car struct {
	Vehicle
	Horsepower int
}
