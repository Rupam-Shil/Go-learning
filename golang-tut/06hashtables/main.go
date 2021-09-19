package main

import "fmt"

func main() {
	var m map[string]string
	fmt.Println(m == nil)
	m = map[string]string{}
	fmt.Println(m == nil)

	n := make(map[string]string, 5)
	fmt.Println(n)

	p := map[string]string{"name":"rupam"}
	p["anotherName"] = "subho"
	fmt.Println(p)

	for j, i := range p{
		fmt.Println(i,j)
	}

	title, ok := p["soumya"]
	if ok{
		fmt.Println(title)
	} else{
		fmt.Println("could find")
	}
}