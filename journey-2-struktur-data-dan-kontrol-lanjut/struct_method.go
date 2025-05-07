package main

import "fmt"

// Mendefinisikan struct
type Person struct {
	Name string
	Age  int
}

// Method untuk struct
func (p Person) introduce() {
	fmt.Println("Hello, my name is", p.Name, "and I am", p.Age, "years old.")
}

func main() {
	person := Person{Name: "Imam", Age: 30}
	person.introduce()
}
