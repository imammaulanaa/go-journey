package main

import "fmt"

func relax() {
	if r := recover(); r != nil {
		fmt.Println("Recover dari panic:", r)
	}
}

func Panic() {
	panic("terjadi panic")
}

func main() {
	defer relax()
	Panic()
	fmt.Println("Program tetap jalan setelah recover")
}
