package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Practice")
	fmt.Println("What's your name?")

	// reader := bufio.NewReader(os.Stdin)
	// name, _ := reader.ReadString('\n')
	// fmt.Println(name)
	time := time.Now()
	fmt.Println(time.Format("Monday"))
	type hola string;
	var myName hola = "hi there"
	fmt.Println(myName)
	const fixedValue = 5;
	if fixedValue > 10{
		fmt.Println("greated")
	} else {
		fmt.Println("Lesser")
	}
	
}