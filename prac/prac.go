package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to practice file 2")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name")
	name, err := reader.ReadString('\n')
	checkError(err)
	fileContent := "hello" + name
	filePath := "./" + strings.TrimSpace(name)  + ".txt"
	file, err := os.Create(filePath)
	checkError(err)
	length, err := io.WriteString(file,fileContent)
	checkError(err)
	fmt.Println("The length of the file",length)
	defer file.Close()
}

func checkError(err error){
	if err != nil {
		panic(err)
	}
}