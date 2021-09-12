package main

import (
	"fmt"
	"math"
)

func main() {
	num := 10
	var num32 int32 = 34	
	var num64 int64 = 64	
	fmt.Println(num)
	fmt.Printf("The number is %v and the type is  %T\n",num,num)
	fmt.Printf("The number is %v and the type is  %T\n",num32,num32)
	fmt.Printf("The number is %v and the type is  %T\n",num64,num64)

	num = int(num32)

	var unsigned uint = 1
	fmt.Println(unsigned)
	//Floating point numbers
	py := math.Pi
	fmt.Println(py)
	fmt.Printf("The number type is %T\n",py)

	ninjas := 1e100
	fmt.Println(ninjas)

	aboutFour := 3.999999
	fmt.Println(math.Floor(aboutFour))

	fmt.Println(complex(5,4))

}