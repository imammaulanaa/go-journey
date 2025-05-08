package main

import (
	"fmt"
	"time"
)

func greet(name string) {
	time.Sleep(2 * time.Second)
	fmt.Println("Hello", name)
}

func main() {
	go greet("Imam")
	go greet("Akbar")

	time.Sleep(3 * time.Second) // Tunggu goroutine selesai
}
