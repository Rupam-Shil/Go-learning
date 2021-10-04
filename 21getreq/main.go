package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to get request")
	PerformGetRequest()
}

func  PerformGetRequest(){
	const myurl string = "http://localhost:8000/get"
	res,err := http.Get(myurl)
	checkError(err)
	defer res.Body.Close()
	fmt.Println("Status Code: ", res.StatusCode)
	fmt.Println("Content Type: ",res.ContentLength)
	content,_ :=ioutil.ReadAll(res.Body)
	
	stringBuilder := strings.Builder{}

	byteCount , _ := stringBuilder.Write(content)
	fmt.Println(byteCount)
	fmt.Println(stringBuilder.String())
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}