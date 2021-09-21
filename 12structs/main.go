package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct in golang")
	//no inheritance :No super or parent

	rupam := User{"Rupam",[]string{"rupam@go.dev"},true,19}
	fmt.Println(rupam.Email)
	fmt.Printf("Rupam's details are %+v", rupam)
}

type User struct{ 
	Name 	string
	Email 	[]string
	Status 	bool
	Age 	int

}