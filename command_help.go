package main

import "fmt"

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	//printCmd()
	// fmt.Println("help: Display a help message")
	// fmt.Println("exit: Exit the Pokedex")
	for name, cmd := range fetchCmd() {
		fmt.Println(name, ": ", cmd.description)
	}
	return nil
}
