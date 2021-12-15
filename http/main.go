package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println(res)
}

type Shape interface {
	area() int
}

type Area struct {
	sileLength int
}

func (a Area) calculate() int {
	return 10
}

func print(s Shape) {
	fmt.Println(s.area())
}
