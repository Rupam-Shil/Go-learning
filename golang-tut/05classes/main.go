package main

import "fmt"

func main() {
	fmt.Println("Welcome to class")


	type ninja struct {
		name string
		weapons []string
		level int
	}

	rupam := ninja{
		name:"Rupam",
	}

	rupam = ninja{
		name:"Rupam",
		weapons: []string{"shord","star"},
		level: 5,
	}
	fmt.Println(rupam)

	rupam = ninja{
		"Subho",
		 []string{"shadow","star"},
		 1,
	}

	fmt.Println(rupam)
	fmt.Printf("the type of weapons is %T\n", rupam.weapons)

	type dojo struct {
		name string
		ninja ninja
	}
	type adressdojo struct {
		name string
		ninja *ninja
	}
	// community1 := dojo{"ninjastars", rupam}
	community2 := adressdojo{"ninjaman", &rupam}
	community2.ninja.level = 8
	fmt.Println(*community2.ninja)
	fmt.Println(rupam.level)

	adhiPointer := new(ninja)
	fmt.Println(adhiPointer)

	intern := ninjaIntern{name:"Intern", level:1}
	intern.sayHello()
}

type ninjaIntern struct{
	name string
	level int
}

func (ninjaIntern) sayHello(){
	fmt.Println("Hello")
}