package main

import (
	"errors"
	"fmt"
	"strings"
)

var ErrEmptyString = errors.New("Received Empty String")

func ToLowerCase(s string) (string, error) {
	if s == "" {
		return "", ErrEmptyString
	}
	return strings.ToLower(s), nil
}

func main() {
	fmt.Println("Hello World")
}
