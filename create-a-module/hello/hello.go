package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set properties of the predefined Logger, including:
	// 1) the log entry prefix
	// 2) a flag to disable printing the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"jaime", "alex", "danny"}

	messages, err := greetings.Hellos(names)

	// if an error was returned, print to console + exit program
	if err != nil {
		log.Fatal(err)
	}
	// otherwise, return the message
	fmt.Println(messages)
}