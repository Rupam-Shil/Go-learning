package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	subho := Student{"subho", 20}
	rupam  := &subho
	fmt.Println(subho.age)
	fmt.Println(*rupam)
}