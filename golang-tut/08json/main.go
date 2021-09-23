package main

import (
	"encoding/json"
	"fmt"
)

type Ninja struct {
	Name   string   `json:"full_name"`
	Weapon Weapon
	Level  int
}

type Weapon struct {
	Name string `json:"weapon_name"`
	Level string `json:"weapon_level"`
}

func main() {
	sFrom := `{"full_name":"Rupam","weapon":{"weapon_name":"ninja star","weapon_level":"4"},"level":1}`
	fmt.Println(sFrom)
	var rupam Ninja;
	err := json.Unmarshal([]byte(sFrom), &rupam)  //for json -> object
	if err != nil {
		fmt.Println("😢")

		fmt.Println(err)
	}
	fmt.Println(rupam)

	sTi, err := json.Marshal(rupam)  // for object -> json
	if err != nil {
		fmt.Println("😢")

	}
	fmt.Println(sTi)  // retuns byte
	fmt.Printf("%s\n", sTi)

	
}