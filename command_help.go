package main

import "fmt"

func commandHelp(c *config, a string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	//printCmd()
	// fmt.Println("help: Display a help message")
	// fmt.Println("exit: Exit the Pokedex")
	for name, cmd := range fetchCmd() {
		fmt.Println(name, ": ", cmd.description)
	}
	return nil
}
