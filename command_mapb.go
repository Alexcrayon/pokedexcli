package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMapb(c *config) error {

	if c.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url := *c.Previous
	fmt.Println("Prev URL:", url)
	var body []byte
	if cached, ok := c.cache.Get(url); ok {
		body = cached
		fmt.Println("map back cache hit")
	} else {
		fmt.Println("map back cache missed")
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		body, err = io.ReadAll(res.Body)
		defer res.Body.Close()
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
