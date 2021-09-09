package main

import (
	"fmt"
	"math/big"

	// "math/rand"

	"crypto/rand"
)

func main() {
	fmt.Println("welcome to math in golang")
	// var myNumberOne int  = 2
	// var myNumbertwo float64 = 4
	// fmt.Println("The sum is:",m)
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))
	randomNum, _ :=  rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randomNum)


}