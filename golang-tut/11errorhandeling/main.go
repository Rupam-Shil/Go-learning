package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Welcome to golnag")
	val,err := returnError(false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}

type specialError struct {
	errorMesage string
	errorCode int
}


func (s specialError) Error() string {
	return s.errorMesage + " " + strconv.Itoa(s.errorCode)
}

func returnError(returnError bool) (string,error) {
	if returnError {
		return "", specialError{"special error",10}
	}

	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(5)
	seedString := strconv.Itoa(randomNum)

	return seedString,nil
}