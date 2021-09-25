package main

import "fmt"

type ninjaStar struct {
	owner string
}

type ninjaSword struct {
	owner string
}

type ninjaWeapon interface {
	attack()
}

func attack(weapon ninjaWeapon){
	weapon.attack()
}

func (s ninjaSword) attack(){
	fmt.Println("Throwing ninja sword",s.owner)
}

func (s ninjaStar) attack() {
	fmt.Println("Throwing ninja star",s.owner)
}

func main() {
	stars := []ninjaStar{{"rupam"}, {"subho"}}

	for _, star := range stars {
		star.attack()
	}
	fmt.Println("")

	swords := []ninjaSword{{"rupam"}, {"subho"}}
	for _, sword := range swords {
		sword.attack()
	}
	fmt.Println("")
}