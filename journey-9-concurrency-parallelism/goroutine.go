package main

import (
    "fmt"
    "time"
)

func cetakPesan(pesan string) {
    for i := 0; i < 3; i++ {
        fmt.Println(pesan)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    go cetakPesan("Hello dari Goroutine")
    cetakPesan("Hello dari Main")
}
