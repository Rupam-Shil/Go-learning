package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file"

	file, err := os.Create("./myText.txt")
	checkNilErr(err)


	length, err :=io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is ", length)
	defer file.Close()
	readFile("./myText.txt")

}


func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Printf("%s", databyte)

}

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}