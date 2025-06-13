# Journey 9: Concurrency & Parallelism
> **Cerita**  
Gopher mengangkat panji Parallelism, mengerahkan legiun Goroutine ke medan laga, mengkoordinasikan serangan lewat Channel bersih. Dengan `select`, ia memilih jalur terbaik di persimpangan tugas, dan membangun Worker Pool seperti markas lapangan untuk menjalankan operasi skala besar.


## Tujuan Pembelajaran
- Spawning goroutine
- Unbuffered & buffered channel
- Select statement
- Worker pool pattern

## Daftar File
- `goroutine.go`
- `channel.go`
- `buffered_channel.go`
- `select.go`
- `worker_pool.go`

## Cara Menjalankan
```bash
cd journey-9-concurrency-parallelism
go run worker_pool.go
```

## Catatan
Gunakan WaitGroup untuk sinkronisasi.

## Referensi
- https://gobyexample.com/goroutines
- https://gobyexample.com/channels