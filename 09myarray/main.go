package main

import "fmt"

func main() {
	fmt.Println("Welcome to array to golang")
	var fruitList [4]string

	fruitList[0]  = "apple"
	fruitList[1]  = "Tomato"
	fruitList[3]  = "Banana"

	fmt.Println(len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "EggPlant"}
	fmt.Println(vegList)

}