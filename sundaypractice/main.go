package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Details struct {
	Name string 
	Usn string 
	Password string 
}

func main() {
	fmt.Println("Welcome to sunday practice")
	performPostRequest()
}

func performPostRequest(){
	const url string = "http://localhost:8000/post"
	rupam := Details{"Rupam","ENG19CS0265","test123"}
	
	finalJson ,err := json.Marshal(rupam)
	checkError(err)

	fmt.Printf("%s", finalJson)
	dataToSend := strings.NewReader(string(finalJson))
	res, err := http.Post(url, "application/json", dataToSend)
	checkError(err)
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Println("here")
	fmt.Println(string(content))
}

func checkError(err error){
	if err != nil{
		panic(err)
	}
}