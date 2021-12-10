package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string // use custom type
type Deck struct {
	Cards []Card // use struct
}

type Card struct {
	Suite string // use struct
	Value string // use struct
}

func (c deck) print() {
	for i, cards := range c {
		fmt.Println(i, cards)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuite := []string{"spades", "hearts", "diamonds", "clubs"}
	cardValues := []string{"One", "Two", "Three", "Ford"}

	for _, suite := range cardSuite {
		for _, value := range cardValues {
			// cards.Cards = append(cards.Cards, Card{Suite: suite, Value: value})
			cards = append(cards, value+" "+suite)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func writeFile(d *Deck, data []byte, perm os.FileMode) error {
	out, err := os.Open("main.go")
	if out != nil {
		arr := [5]int{0, 0, 0, 0, 0}
		fmt.Println(arr)
		fmt.Println("open file", out)
	}
	return err
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0644)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName) // bs is abbreviation of byteSlice
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), "")
	return deck(s)
}

func (d deck) shuffle() {
	var source = rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[i], d[newPosition]
	}
}
