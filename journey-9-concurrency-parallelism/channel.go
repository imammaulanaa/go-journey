package main

import "fmt"

func kirimData(ch chan string) {
    ch <- "data dari goroutine"
}

func main() {
    ch := make(chan string)
    go kirimData(ch)
    pesan := <-ch
    fmt.Println("Menerima:", pesan)
}
