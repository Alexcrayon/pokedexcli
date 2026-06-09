package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

type location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"` // pointer: null at the last page
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		//URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(c *config) error {

	url := "https://pokeapi.co/api/v2/location-area/"
	if c.Next != nil {
		url = *c.Next
	}
	pokecache.NewCache(10 * time.Second)

	//} else {
	//	fmt.Println("you're at last page")
	//	return nil
	//}

	var body []byte
	if cached, ok := c.cache.Get(url); ok {
		body = cached
	} else {
		res, err := http.Get(url)
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

	var loc location
	if err := json.Unmarshal(body, &loc); err != nil {
		return err
	}

	c.Next = loc.Next
	c.Previous = loc.Previous
	for _, result := range loc.Results {
		fmt.Println(result.Name)
	}
	//fmt.Printf(loc.Results[0].Name)
	return nil
}
