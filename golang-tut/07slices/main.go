package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices ")
	fixedSizedArray := [...]string{"js", "rb", "py"}
	fmt.Printf("The type of fixedSizedArray is %T\n", fixedSizedArray)

	arrSlice := make([]string, 2)
	arrSlice = []string{"js", "rb", "bola"}
	arrSlice = append(arrSlice, "py")
	var index int  = 2;
	arrSlice = append(arrSlice[:index], arrSlice[index+1:]...)
	fmt.Println(arrSlice)
	fmt.Printf("The type of fixedSizedArray is %T\n", arrSlice)
	fmt.Println(len(arrSlice), cap(arrSlice))

}