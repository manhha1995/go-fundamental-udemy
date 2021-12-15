package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Square struct {
	sideLength float64
}

type Rectangle struct {
	base   float64
	height float64
}

func (a Square) getArea() float64 {
	return a.sideLength * a.sideLength
}

func (b Rectangle) getArea() float64 {
	return b.height * b.base * 0.5
}

func print(s Shape) {
	fmt.Println(s.getArea())
	fmt.Println(s.getArea())
}

func main() {
	a := Rectangle{3.5, 3.5}
	b := Square{3.5}
	print(a)
	print(b)
}
