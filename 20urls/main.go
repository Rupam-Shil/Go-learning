package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://www.youtube.com/watch?v=cl7_ouTMFh0&ab_channel=HiteshChoudhary"

func main() {
	fmt.Println("Welcome to handling URL in golang")
	fmt.Println(myUrl)

	// parsing 
	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()

	for key, value := range qparams {
		 fmt.Println(key, value)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tut",
		RawQuery: "user=hitesh",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}