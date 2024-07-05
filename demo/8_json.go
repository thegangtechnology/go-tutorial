package demo

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"` //these are called struct tag
	Age  int    `json:"age"`  //case insensitive (i believe)
}

func parsePerson(s string) (Person, error) {
	var person Person
	err := json.Unmarshal([]byte(s), &person)
	if err != nil {
		return Person{}, fmt.Errorf("parse error: %w", err)
	}
	return person, nil
}
