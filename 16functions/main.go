package main

import "fmt"

type Rectangle struct {
	height float64
	width float64
}


func main() {
	fmt.Println("Welcome to functions")

	a1 := Rectangle{4,5}
	fmt.Println(a1.area())
	greeterThree(greeterTwo)()
	total,message := proAdder(5,6,7,8,9,10,11,12,13,14,15)
	fmt.Println(message,total)
}

func greeterTwo(){
	fmt.Println("first")
}

func greeterThree(f func()) func(){
	return f
}

func proAdder(values ...int)(int,string){
	total := 0
	for _,  value := range values {
		total += value
	} 
	return total, "The total value is"
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}