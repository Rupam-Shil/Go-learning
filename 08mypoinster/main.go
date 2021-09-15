package main

import "fmt"

func main() {
	var num int = 20
	var ptr *int = &num
	fmt.Printf("The value is %v\n", *ptr)
	fmt.Printf("The value is %v\n", ptr)
	*ptr = *ptr + 20
	fmt.Printf("The value is %v\n", *ptr)
	anotherPtr := ptr
	fmt.Printf("The value is %v\n", *anotherPtr)


}