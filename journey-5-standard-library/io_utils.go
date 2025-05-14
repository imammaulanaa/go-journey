package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Halo Imam!\n")
	ioutil.WriteFile("contoh.txt", data, 0644)

	content, err := ioutil.ReadFile("contoh.txt")
	if err != nil {
		fmt.Println("Gagal membaca file:", err)
		return
	}

	fmt.Println("Isi file:", string(content))

	os.Remove("contoh.txt") // Bersihkan file
}
