package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message, err := greetings.Hello("Lady")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"abby", "ben", "clare"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
