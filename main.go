package main

import (
	"fmt"
	"go-lessons-alua/greet"
	"go-lessons-alua/greet/input"
)

func main() {
	var userInput input.Input

	fmt.Print("Name: ")
	_, err := fmt.Scanln(&userInput.Name)
	if err != nil {
		fmt.Println("error while reading name from input: ", err)
	}

	fmt.Print("Table: ")
	_, err = fmt.Scanln(&userInput.Table)
	if err != nil {
		fmt.Println("error while reading table from input: ", err)
	}

	output := greet.Greet(userInput)
	fmt.Println(output)
}
