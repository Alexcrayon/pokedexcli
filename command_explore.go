package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type locPokemon struct {
	//Count    int     `json:"count"`
	//Next     *string `json:"next"` // pointer: null at the last page
	//Previous *string `json:"previous"`
	PokemonEncounter []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`

		//URL  string `json:"url"`
	} `json:"pokemon_encounters"`
}

func commandExplore(c *config, name string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + name

	//if c.Next != nil {
	//	url = *c.Next
	//}
	var body []byte
	if cached, ok := c.cache.Get(url); ok {
		body = cached
		//fmt.Println("explore Cache Hit")
	} else {
		//fmt.Println("explore Cache missed")
		res, err := http.Get(url)
		fmt.Println("Exploring " + name + "...")
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)

		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}

		c.cache.Add(url, body)
	}

	var locPoke locPokemon
	if err := json.Unmarshal(body, &locPoke); err != nil {
		return err
	}

	//c.Next = loc.Next
	//c.Previous = loc.Previous
	fmt.Println("Found Pokemon:")
	for _, result := range locPoke.PokemonEncounter {
		fmt.Println(" - " + result.Pokemon.Name)
	}
	//fmt.Printf(loc.Results[0].Name)

	return nil
}
