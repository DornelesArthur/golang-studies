package main

import (
	"fmt"
)

/*
bool

string

int Both int and uint contain same size, either 32 or 64 bit.
int8		8-bit signed integer
int16		16-bit signed integer
int32		32-bit signed integer
int64		64-bit signed integer

uint 		Both int and uint contain same size, either 32 or 64 bit.
uint8		8-bit unsigned integer
uint16		16-bit unsigned integer
uint32		32-bit unsigned integer
uint64		64-bit unsigned integer
uintptr		It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.

byte		It is a synonym of uint8.

rune		It is a synonym of int32 and also represent Unicode code points.

float32		32-bit IEEE 754 floating-point number
float64		64-bit IEEE 754 floating-point number

complex64	Complex numbers which contain float32 as a real and imaginary component.
complex128	Complex numbers which contain float64 as a real and imaginary component.
*/

func main() {
	var number uint8
	// numero = -1
	// number = 256
	number = 255

	fmt.Println(number)

	var complex_number complex64
	complex_number = 255 + 11i

	fmt.Println(complex_number)
}