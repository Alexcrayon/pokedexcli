package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//fmt.Println("Hello, World!")
	scanner := bufio.NewScanner(os.Stdin)

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

		//fmt.Println("Your command was:" + cleanInput(in)[0])
		fmt.Println("Your command was: " + cleanInput(in.String())[0])
	}

}

func cleanInput(text string) []string {
	//tolower case first then clean the space
	lower := strings.ToLower(text)
	cleaned := strings.Fields(lower)
	return cleaned
	//return nil
}
