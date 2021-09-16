package main

import "fmt"

func main() {
	for i := 0; i < 10; i+=2 {
		fmt.Println("Hello")
	}

	var s string = "Hello world"
	for i :=0 ; i < len(s) ; i++ {
		fmt.Printf("%v %c\n",i, s[i])
	}

	for i, c:= range s {
		fmt.Printf("%d %c", i, c)
		fmt.Println("")
	}

	//break
	for i, c:= range s {
		
		if i ==4 {
			break
		}
		fmt.Printf("%d %c", i, c)
		fmt.Println("")
	}

	//label
	iForLoop:
	for i:=0;i<5;i++{
		for j :=0; j <5; j++ {
			if j ==3{
				break iForLoop  //breaks the i for loop because of the label
			}

			fmt.Printf("%d%d\n", i,j)

		}
		fmt.Println("")
	}
}