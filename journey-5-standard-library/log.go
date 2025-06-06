package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println("Aplikasi dimulai...")

	log.Println("Menjalankan sesuatu...")

	log.Println("Aplikasi selesai.")
}
