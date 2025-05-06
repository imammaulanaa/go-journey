package main

import "fmt"

func main() {
    // if-else
    nilai := 75
    if nilai >= 80 {
        fmt.Println("Lulus dengan nilai bagus")
    } else if nilai >= 60 {
        fmt.Println("Lulus")
    } else {
        fmt.Println("Tidak lulus")
    }

    // switch
    hari := "Senin"
    switch hari {
    case "Senin":
        fmt.Println("Hari kerja")
    case "Sabtu", "Minggu":
        fmt.Println("Akhir pekan")
    default:
        fmt.Println("Hari biasa")
    }

    // for loop
    for i := 1; i <= 5; i++ {
        fmt.Println("Iterasi ke:", i)
    }
}