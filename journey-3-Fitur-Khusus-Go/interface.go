package main

import "fmt"

// Definisi interface
type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

func main() {
	var speaker Speaker = Person{Name: "Imam"}
	fmt.Println(speaker.Speak())
}