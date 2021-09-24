package main

import (
	"fmt"
)

func main() {
	// fmt.Println(sum(2,3)) 
	// greeting("Welcome")
	// vardict(12,3,6,5)
	fmt.Println(nameLength("Rupam"))
	fmt.Println(fibonacci(10))
	result, ok := greetingOk("sdfsd")
	if !ok  {
		fmt.Println(ok)
	} else {
		fmt.Println(result)
	}
}

// func greeting(value string){
// 	fmt.Println(value)
// }

// func vardict(intergers ...int){
// 	for _,interger := range intergers {
// 		fmt.Println(interger)
// 	}
// }

// func sum(a, b int) int{
// 	return a+b
// }

func nameLength(name string) (string, int){
	return name, len(name)
}

func greetingOk(value string) (string, bool){
	if len(value) == 0 {
		return "", false
	} else {
		return "Hello" + value, true
	}
}

func fibonacci(i int) int{
	if i ==0{
		return 0
	} else if i == 1{
		return 1
	} else {
		return fibonacci(i-1) + fibonacci(i-2)
	}
}