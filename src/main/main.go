package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Card struct {
	Name   string `json:"name"`
	Life   int    `json:"life"`
	Attack int    `json:"attack"`
}

func main() {
	cards := getCards()

	for _, value := range cards {
		fmt.Println(value)
	}
}

func getCards() []Card {
	res, err := ioutil.ReadFile("./data/example.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var c []Card
	json.Unmarshal(res, &c)
	return c
}
