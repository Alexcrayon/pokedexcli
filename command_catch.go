package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type CatchPokemon struct {
	BaseExp string `json:"base_experience"`
}

func commandCatch(c *config, a string) error {
	fmt.Println("Throwing a Pokeball at " + a + "...")

	url := "https://pokeapi.co/api/v2/pokemon/" + a
	var body []byte
	if cached, ok := c.cache.Get(url); ok {
		body = cached
		//fmt.Println("explore Cache Hit")
	} else {
		//fmt.Println("explore Cache missed")
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

	var pokemon CatchPokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return err
	}

	catch(pokemon.BaseExp)

	return nil
}

func catch(exp string) float64 {
	base, _ := strconv.Atoi(exp)
	prob := float64(base / 1000)
	//need flip the probability
	return prob
}
