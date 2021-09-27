package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to methods")
	rupam := User{"Rupam", "rupamshil111@gmail.com", true, 21,70}
	rupam.GetStatus()
	rupam.NewMail()
	fmt.Println(rupam.Email)

}

type User struct {
	Name string
	Email string
	Status bool
	Age int
	oneAge int
}

func (u User) GetStatus(){
	fmt.Println(u.Status)
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of the user is: ", u.Email)
}