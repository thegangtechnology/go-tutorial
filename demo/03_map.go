package demo

import "fmt"

func MinKeyValue(m map[string]int) (string, error) {
	if len(m) == 0 {
		return "", fmt.Errorf("map is empty")
	}
	minKey := "" //note first assignment
	minValue := 0
	for k, v := range m {
		if v < minValue {
			minValue = v
			minKey = k // and reassignment
		}
	}
	return minKey, nil
}

func CharCount(s string) map[string]int {
	var words map[string]int

	//words := make(map[string]int) //same here
	for _, word := range s {
		sw := string(word)
		v, ok := words[sw]
		if ok {
			words[sw] = v + 1
		} else {
			words[sw] = 1
		}
	}
	return words
}
