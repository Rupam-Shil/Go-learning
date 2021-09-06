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
	// var godKnows []string
	godKnows := [3]string{"rupam", "subho", "randomname"}
	r := []bool{true, false, true, true, false, true}
	for i:=0; i
}