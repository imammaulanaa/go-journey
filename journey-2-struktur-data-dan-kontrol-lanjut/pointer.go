package main

import "fmt"

func main() {
	x := 10
	ptr := &x // Pointer ke x

	fmt.Println("Nilai x:", x)
	fmt.Println("Alamat memori x:", ptr)