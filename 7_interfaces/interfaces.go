package main

import (
	"fmt"
	"math"
)

type geometric interface {
	area() float64
	perimeter() float64
}

type square struct {
	height int
}

func (s square) area() float64 {
	return float64(s.height * s.height)
}

func (s square) perimeter() float64 {
	return float64(s.height * 4)
}

type circle struct {
	radius int
}

func (c circle) area() float64 {
	return math.Pi * float64(c.radius) * float64(c.radius)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * float64(c.radius)
}

type triangle struct {
	height int
	length int
	a      int
	b      int
	c      int
}

func (t triangle) area() float64 {
	return float64(t.height*t.length) / 2
}

func (t triangle) perimeter() float64 {
	return float64(t.a + t.b + t.c)
}

func ShowArea(g geometric) {
	fmt.Println(g.area())
}

func ShowPerimeter(g geometric) {
	fmt.Println(g.perimeter())
}

func main() {
	t := triangle{
		height: 3,
		length: 4,
		a:      3,
		b:      4,
		c:      5,
	}
	ShowArea(t)
	ShowPerimeter(t)

	c := circle{
		radius: 3,
	}
	ShowArea(c)
	ShowPerimeter(c)

	s := square{
		height: 3,
	}
	ShowArea(s)
	ShowPerimeter(s)

	fmt.Println("==========================================")

	var list []interface{}
	list = append(list, 2)
	list = append(list, true)
	list = append(list, "string")
	list = append(list, 7.5)
	list = append(list, t)
	fmt.Println(list)

	for _, value := range list {
		if v, ok := value.(string); ok {
			fmt.Println(v + " is a string")
		} else {
			fmt.Println(value)
		}
	}
}
