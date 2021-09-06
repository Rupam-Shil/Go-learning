package main

import "fmt"

func main() {
	username := []string{"hello", "dear"}
	fmt.Println(username)
	fmt.Printf("The variable is of type %T \n", username)
	var smallFloat float64 = 4.444444444444444444444444
	fmt.Println(smallFloat)
	fmt.Printf("The variable is of type %T \n", smallFloat)

	//default and aliases
	var anotherVariable int;
	fmt.Println(anotherVariable);
	fmt.Println("The variable is of type %T \n", anotherVariable);

	var anotherString string;
	fmt.Println(anotherString);
	fmt.Println("The variable is of type %T \n", anotherString);

	var website = "rupam@gmail.com";
	fmt.Println(website);

	//no var type
	numberOfUser := 300000
	fmt.Println(numberOfUser)

}