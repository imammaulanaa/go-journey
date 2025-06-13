# Journey 7: Error Handling & Custom Error
> **Cerita**  
Saat malam pekat, Gopher menenun jaring error, menciptakan Custom Error Types untuk menangkap kesalahan dengan rapi. Ia membungkus pesan dengan `fmt.Errorf` dan belajar membangkitkan `panic`—namun selalu siap menghidupkan kembali istana dengan `recover` sebelum segalanya runtuh.


## Tujuan Pembelajaran
- Error wrapping dengan %w
- Custom error type
- Panic & recover

## Daftar File
- `error_basic.go`
- `error_wrapping.go`
- `custom_error.go`
- `panic_recover.go`

## Cara Menjalankan
```bash
cd journey-7-error-handling
go run error_basic.go
```

## Catatan
Gunakan %w untuk unwrap.

## Referensi
- https://blog.golang.org/errors-are-values
- https://pkg.go.dev/errors