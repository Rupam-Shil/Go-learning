package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //usually these are pointers

func main() {
	// go greeter("Hello")
	// greeter("World")
	websiteList := []string{"http://lco.dev","http://go.dev","http://google.com", "https://github.com", "https://amazon.com"}
	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

// func greeter(s string) {
// 	for i :=0; i <6;i++{
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string){
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil{
		fmt.Println("Upps in endpoint")
	}
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}