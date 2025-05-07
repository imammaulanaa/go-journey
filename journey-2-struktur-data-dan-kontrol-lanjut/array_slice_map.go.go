package main

import "fmt"

func main() {
	// Array
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Slice
	slice := []int{10, 20, 30}
	slice = append(slice, 40)
	fmt.Println("Slice:", slice)

	// Map
	student := map[string]int{
		"Alice": 90,
		"Bob":   85,
	}
	fmt.Println("Student Map:", student)
}
