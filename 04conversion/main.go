package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to typecasting")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a rating between 1 and 5")
	rating , _ := reader.ReadString('\n')

	updatedRating , err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The updated rating is %v", updatedRating+1)
	}
}