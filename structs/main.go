package main

import "fmt"

type person struct {
	name        string
	phoneNumber int
}

func main() {
	characters := person{name: "John", phoneNumber: 1234}
	newPointer := &characters
	newPointer.updateName("Jennifer")
	fmt.Printf("%+v", newPointer)
}
func (ptr *person) updateName(newName string) {
	(*ptr).name = newName
}
