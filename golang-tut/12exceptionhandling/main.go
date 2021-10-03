package main

import "fmt"

func main() {
	defer safeExit()
	fmt.Println("Welcome to panic and recover")
	const one int = 2;
	if one != 1{
		panic("One is not 1")
	}

	

}

func safeExit(){
	if r := recover(); r != nil {
			fmt.Println("Panic is recovered")
		}
}