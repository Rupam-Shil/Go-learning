package main

import "fmt"

func main() {
	fmt.Println("Welcome to function part 2")
	fmt.Printf("The type is %T\n",returnFunction(9, returnParam))

	//defer first()
	second()
	second()

	aFunctinon := func(){
		fmt.Println("anomonous")
	}
	aFunctinon()

	// aFunctinon = first     -> possible as aFunction and first are of same type
	// aFunctinon = returnParam   -> not possible as aFunction and returnParam are not of same type

	
	crazyFunction := func(f func()) func(){
		return  f
	}
	crazy(crazyFunction)(first)()
}

func crazy(f func(func()) func()) func(func()) func(){
	return f
}


func first() {
	fmt.Println("First")
}

func second(){
	fmt.Println("Second")
}

func returnParam(i int) int {
	return i
}

func returnFunction(i int, f func(int)int) int {
	return f(i)
}