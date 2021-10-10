package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Welcome to golang practice")
	const url string = "https://api.nuxtjs.dev/mountains"
	var jsonArray []interface{}
	res, err := http.Get(url)
	panicError(err)
	defer res.Body.Close()
	content , _:= ioutil.ReadAll(res.Body)
	json.Unmarshal(content, &jsonArray)
	fmt.Printf("%#v\n",jsonArray)
	for _, value := range jsonArray {
		fmt.Println(value)
	}
}

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}