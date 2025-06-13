# Journey 17: Advanced Error Handling
> **Cerita**  
Gopher mengasah pedang Error Handling, membungkus kesalahan dengan lapisan `fmt.Errorf` dan `%w`, lalu menelisik akar masalah melalui `errors.Unwrap`. Ia menciptakan sentinel error, memastikan setiap jenis kegagalan dikenali dan ditangani dengan bijak.


## Tujuan Pembelajaran
- Error wrapping & unwrapping
- Custom error struct
- errors.Is & errors.As

## Daftar File
- `customerror/error.go`
- `utils/wrap.go`
- `handler/service.go`
- `main.go`

## Cara Menjalankan
```bash
cd journey-17-error-handling-advanced
go run main.go
```

## Catatan
Gunakan fmt.Errorf dengan %w.

## Referensi
- https://pkg.go.dev/errors