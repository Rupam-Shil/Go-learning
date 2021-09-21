package main

import "fmt"

func main() {
	venues := []string{"Home", "Work", "School", "Bar", "Gym","check"}

	for _, venue := range venues {
		switch venue {
		case "Home":
			greeting("In home")
			break
		case "Work":
			greeting("I am at work! Can't talk")
			break
		case "School":
			greeting("I hate my school")
			break
		case "Bar":
			greeting("Feeling dizzy")
			break
		case "Gym":
			greeting("One more set")
			break
		default:
			greeting("huh")
		}
	}
}

func greeting(greeting string) {
	fmt.Println(greeting)
}