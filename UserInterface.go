package main

import (
	"log"
	"github.com/manifoldco/promptui"
	"fmt"
	"strconv"
	"errors"
	"strings"
)

func yes_no_prompt(label string) bool {
    prompt := promptui.Select{
        Label: "string",
        Items: []string{"Yes", "No"},
    }
    _, result, err := prompt.Run()
    if err != nil {
        log.Fatalf("Prompt failed %v\n", err)
    }
    return result == "Yes"
}

func number_prompt() (int, error) {
	var input string
	var number int

	fmt.Scan(&input)
	if set_contains([]string{"exit", "quit", "q"}, strings.ToLower(strings.TrimSpace(input))) {
		return -1, nil
	}
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("\"%s\" Could not be converted into an int: (%v) ", input, err))
	}
	if number <= 0 {
		return 0, errors.New(fmt.Sprintf("\"%d\" is not a positive integer", number))
	}
	return number, nil
}


func string_prompt() (string) {
	var input string
	var output string

	fmt.Scan(&input)
	output = strings.ToLower(strings.TrimSpace(input))
	
	if set_contains([]string{"exit", "quit", "q"}, output) {
		return "exit"
	}
	return output
}

func set_contains(set []string, input string) bool {
	for _, element := range set {
		if element == input {
			return true
		}
	}
	return false
}
