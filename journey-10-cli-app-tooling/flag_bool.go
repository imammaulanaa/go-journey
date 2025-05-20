package main

import (
    "flag"
    "fmt"
)

func main() {
    verbose := flag.Bool("v", false, "Tampilkan log verbose")
    flag.Parse()

    if *verbose {
        fmt.Println("Mode verbose aktif")
    } else {
        fmt.Println("Mode normal")
    }
}
