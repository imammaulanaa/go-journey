# Journey 5: Standard Library
> **Cerita**  
Gopher meluncur di atas perahu `net/http`, menelusuri sungai permintaan dan respons. Ia menjelajahi hutan file system dengan `io/ioutil` dan `os`, memetik buah manipulasi string dan konversi data. Waktu mengalir di dataran `time`, sementara bayangan goroutine berkelindan di jembatan `sync`.


## Tujuan Pembelajaran
- HTTP Server & Client
- File I/O (io/ioutil, os)
- Manipulasi String & Konversi
- Time formatting & sleep
- JSON encode/decode & Logging

## Daftar File
- `http.go`
- `io_utils.go`
- `strings.go`
- `time.go`
- `sync.go`
- `json.go`
- `log.go`

## Cara Menjalankan
```bash
cd journey-5-standard-library
go run http.go
```

## Catatan
Set header Content-Type saat JSON.

## Referensi
- https://pkg.go.dev/net/http
- https://pkg.go.dev/encoding/json