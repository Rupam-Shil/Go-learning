package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to get request")
	PerfromPostFormRequest()
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

func PerfromPostRequest(){
	const myUrl string = "http://localhost:8000/post"
	//fake json payload

	resBody := strings.NewReader(`
	{"courseName":"Let's go with golang",
	"platform":"learncodeonline.in"
	}
	`)

	res, err := http.Post(myUrl, "application/json", resBody)
	checkError(err)
	defer res.Body.Close()
	fmt.Println("The status code is:", res.StatusCode)
	fmt.Println("The content length is:",res.ContentLength)

	content, _ := ioutil.ReadAll(res.Body)
	strignBuilder := strings.Builder{}
	byteLength,_ := strignBuilder.Write(content)
	fmt.Println(byteLength)
	fmt.Println(strignBuilder.String())

}

func PerfromPostFormRequest(){
	const myUrl string = "http://localhost:8000/postform"
	//form data
	data := url.Values{}
	data.Add("firstname", "rupam")
	data.Add("lastname", "shil")
	data.Add("email", "rupam@gmail.com")

	res, err := http.PostForm(myUrl, data)
	checkError(err)
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	stringBuilder := strings.Builder{}
	contentLength, err := stringBuilder.Write(content)
	checkError(err)
	fmt.Println("The content length is ", contentLength)
	fmt.Println("THe html is")
	fmt.Println(stringBuilder.String())
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}