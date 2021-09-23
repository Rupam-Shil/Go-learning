package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	switch diceNumber{
		case 1:
			print("You got one")
		case 2:
			print("You got two")
			fallthrough
		case 3:
			print("You got three")
		case 4:
			print("You got four")
		case 5:
			print("You got five")
		case 6:
			print("You got six")
		default:
			print("hui")
	}

}

func print(msg string) {
	fmt.Println(msg)
}