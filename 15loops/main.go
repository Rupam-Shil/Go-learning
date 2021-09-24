package main

import "fmt"

func main() {
	fmt.Println("Welcome to loop")
	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri"}
	for _, day := range days{
	fmt.Printf("Today is %v\n", day)	
	}

	rougueValue := 1

	for rougueValue <10 {
		rougueValue++

		fmt.Println(rougueValue)
		if rougueValue == 2 {
			goto printMe
		}
		if rougueValue == 5 {
			break
		}

		printMe:
		fmt.Println("Print Me")
	}
}