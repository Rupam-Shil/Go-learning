package main

import (
	"fmt"
	"math"
)

func main() {
	const check = 5
	fmt.Println(check)

	fmt.Println(math.Pi)

	const (
		a =3
		b
		c=6
		d
	)
	fmt.Println(a, b, c, d)

	const (
		zero int = iota
		one
		two

	)

	fmt.Println(zero, one, two)
}