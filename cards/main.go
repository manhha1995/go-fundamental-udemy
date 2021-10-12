package main

import (
	"fmt"
)

func main() {
	cards := []string{"One", Cards()}
	cards = append(cards, "Two")
	for _, c := range cards {
	fmt.Println(c)
	}
}

func Cards() string {
	return "Joker"
}