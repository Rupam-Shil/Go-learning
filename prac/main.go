package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to url requests")
	// getUrl := "http://localhost:8000/get"
	// performGetRequest(getUrl)
	postUrl := "http://localhost:8000/post"
	performPostRequest(postUrl)
}

func performGetRequest(url string){
	res, err := http.Get(url)
	checkError(err)
	defer res.Body.Close()
	fmt.Println("The status of the request is:", res.StatusCode)
	fmt.Println("The length of the request is:", res.ContentLength)
	content, err := ioutil.ReadAll(res.Body)
	checkError(err)
	value := getString(content)
	fmt.Println(value)
}

func performPostRequest(url string){
	data := strings.NewReader(`
	{
		"name": "rupam",
		"course":"golang"
	}
	`)
	res, err :=http.Post(url, "application/json",data)
	checkError(err)
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	value := getString(content)
	fmt.Println(value)
	
}

func checkError ( err error ){
	if err != nil {
		panic(err)
	}
}

func getString(content []byte)(string){
	stringBuilder := strings.Builder{}
	_, err := stringBuilder.Write(content)
	checkError(err)
	return stringBuilder.String()
}