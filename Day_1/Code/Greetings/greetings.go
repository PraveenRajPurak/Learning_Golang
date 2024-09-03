package greetings

import (
	"errors"
	"fmt"
)

func Greetings(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}
	message := fmt.Sprintf("Hello, %v", name)
	return message, nil
}
