package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 4, 6, 5}
	fmt.Printf("The array is of type %T\n", a)

	for _, i := range a {
		fmt.Println(i)
	}

	one := &a;

	fmt.Println(one)
}