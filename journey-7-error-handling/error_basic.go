package main

import (
	"errors"
	"fmt"
)

func bagi(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("tidak bisa dibagi dengan nol")
	}
	return a / b, nil
}

func main() {
	hasil, err := bagi(10, 0)
	if err != nil {
		fmt.Println("Terjadi error:", err)
		return
	}
	fmt.Println("Hasil:", hasil)
}
