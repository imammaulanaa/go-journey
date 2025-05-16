package main

import (
	"errors"
	"fmt"
)

func prosesData(data string) error {
	if data == "" {
		return fmt.Errorf("data kosong: %w", errors.New("input tidak valid"))
	}
	return nil
}

func main() {
	err := prosesData("")
	if err != nil {
		fmt.Println("Gagal:", err)
	}
}
