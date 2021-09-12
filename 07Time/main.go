package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time studey of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006  15:04:05 Monday"))

	createdDate := time.Date(2000, time.August,10,25,23,0,0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006  15:04:05 Monday"))
}