package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// returns a greeting for the named person
func Hello(name string) (string, error) {
	// error handling (for if 'name' arg is an empty string)
	if name == "" {
		return "", errors.New("'name' is an empty string")
	}

	message := fmt.Sprintf(generateRandomGreeting(), name)
	return message, nil // 'nil' means no error
}

// returns a map that associates each name with a greeting
func Hellos(names []string) (map[string]string, error) {
	// create a map to associate names with messages
	messages := make(map[string]string)

	// loop through 'names' slice and call 'Hello' func to get a message for each name
	// for (idx, copy of element's value)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

// returns a random greeting from a set of greetings
func generateRandomGreeting() string {
	// create a slice of greetings
	greetings := []string{
		" hi, %v. welcome!",
		" sup, %v?",
		" well, well, well. if it isn't %v.",
	}

	// generate a random index within the range of the length of 'greetings'
	return greetings[rand.Intn(len(greetings))]
}