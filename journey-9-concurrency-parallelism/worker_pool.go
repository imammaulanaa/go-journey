package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range jobs {
        fmt.Printf("Worker %d memproses job %d\n", id, j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {
    const jumlahPekerja = 3
    const jumlahJob = 5

    jobs := make(chan int, jumlahJob)
    results := make(chan int, jumlahJob)
    var wg sync.WaitGroup

    for i := 1; i <= jumlahPekerja; i++ {
        wg.Add(1)
        go worker(i, jobs, results, &wg)
    }

    for j := 1; j <= jumlahJob; j++ {
        jobs <- j
    }
    close(jobs)

    wg.Wait()
    close(results)

    for res := range results {
        fmt.Println("Hasil:", res)
    }
}
