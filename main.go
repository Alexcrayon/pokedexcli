package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokecache"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}
type config struct {
	Next     *string
	Previous *string
	cache    *pokecache.Cache
}

//var cmdMap =

func main() {

	//fmt.Println("Hello, World!")
	scanner := bufio.NewScanner(os.Stdin)
	c := config{
		nil,
		nil,
		pokecache.NewCache(10 * time.Second),
	}
	for {
		fmt.Print("Pokedex > ")
		var in strings.Builder
		for scanner.Scan() {
			next := scanner.Text()
			in.WriteString(next)
			break
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("reading error:", err)
		}

		cmd := cleanInput(in.String())[0]
		if val, ok := fetchCmd()[cmd]; ok {

			err := val.callback(&c)

			if err != nil {
				fmt.Println(err) // see what actually went wrong
				continue         // keep looping
			}
		} else {
			fmt.Println("Unknown command")
		}
		//fmt.Println("Your command was:" + cleanInput(in)[0])
		//fmt.Println("Your command was: " + cleanInput(in.String())[0])
	}

}

func cleanInput(text string) []string {
	//tolower case first then clean the space
	lower := strings.ToLower(text)
	cleaned := strings.Fields(lower)
	return cleaned
	//return nil
}

func fetchCmd() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "List the next 20 location-areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous 20 location-areas",
			callback:    commandMapb,
		},
	}
}
