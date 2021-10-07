package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name string  `json:"coursename"`
	Price int 	
	Platform string `json:"website"`
	Password string  `json:"-"`   //removes the password
	Tags []string   `json:"tags,omitempty"` 
}

func main() {
	fmt.Println("Let's learn about JSON more")
	EncodeJSON()
}

func EncodeJSON(){
	someCourses := []Course{
		{"ReactJS Bootcamp",299,"LCO","abc123",[]string{"web-dev","js"}},
		{"VueJS Bootcamp",599,"AWS","pgq789",[]string{"web-dev","js","vuex"}},
		{"Angular Bootcamp",999,"Coursera","rup435",nil},
	}
	//package this data as json data
	finalJson, err := json.MarshalIndent(someCourses,"","\t")
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n",finalJson)
}