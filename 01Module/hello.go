package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)



func main() {
	log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request a greeting message.
	
    message, err := greetings.Hello("Gladys")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}