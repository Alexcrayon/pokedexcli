package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func commandMap(c *config, a string) error {

	url := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	if c.Next != nil {
		url = *c.Next
	}
	var body []byte
	if cached, ok := c.cache.Get(url); ok {
		body = cached
		fmt.Println("map Cache Hit")
	} else {
		fmt.Println("map Cache missed")
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
