package main

import (
    "flag"
    "fmt"
)

func main() {
    nama := flag.String("nama", "Imam", "Nama pengguna")
    umur := flag.Int("umur", 30, "Umur pengguna")

    flag.Parse()

    fmt.Printf("Halo %s, umur kamu %d tahun\n", *nama, *umur)
}
