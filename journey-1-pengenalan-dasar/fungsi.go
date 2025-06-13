package main

import "fmt"

func sapa(nama string) string {
    return "Halo, " + nama
}

func tambah(a int, b int) int {
    return a + b
}

func bagiDanSisa(a, b int) (int, int) {
    return a / b, a % b
}

func main() {
    fmt.Println(sapa("imam"))
    fmt.Println("10 + 5 =", tambah(10, 5))

    hasil, sisa := bagiDanSisa(13, 4)
    fmt.Println("13 / 4 =", hasil, "sisa", sisa)
} 