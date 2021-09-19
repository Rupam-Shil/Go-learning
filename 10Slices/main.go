package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"mango","orange","banana"}
	fruitList = append(fruitList, "peach", "watermelon")
	fmt.Printf("the type of fruitlist is  %T\n", fruitList)

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 130
	highScore[2] = 345
	highScore[3] = 999

	highScore = append(highScore, 900,33,222)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)

	//how to remove a value from a slices based on index 

	var courses = []string{"vue","reactjs","ruby","python","angular"}

	fmt.Println(courses)

	index := 2;
	courses = append(courses[:index] ,courses[index+1:]...)
	fmt.Println(courses)}