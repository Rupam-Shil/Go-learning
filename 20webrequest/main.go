package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URl = "https://lco.dev/"

func main() {
	fmt.Println("First web request in Golang")

	res, err := http.Get(URl)
	errorCheck(err)

	fmt.Printf("%T\n",res)
	fmt.Printf("%s",*res)

	defer res.Body.Close()  // calles's responsibilities to close the connection
	databytes ,err := ioutil.ReadAll(res.Body)
	errorCheck(err)
	content := string(databytes)
	fmt.Println(content)
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}