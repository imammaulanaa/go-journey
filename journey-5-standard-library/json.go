package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Encode
	p := Person{Name: "Imam", Age: 27}
	jsonData, _ := json.Marshal(p)
	fmt.Println("JSON:", string(jsonData))

	// Decode
	data := `{"name":"Akbar","age":30}`
	var p2 Person
	json.Unmarshal([]byte(data), &p2)
	fmt.Printf("Decoded: %+v\n", p2)
}
