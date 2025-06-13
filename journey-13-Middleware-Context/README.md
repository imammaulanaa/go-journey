# Journey 13: Middleware & Context
> **Cerita**  
Gopher memasang portal Middleware, memeriksa token magis di header, dan menanam konteks dalam tiap permintaan. Ia belajar menyisipkan data user ke dalam arus eksekusi via `context.WithValue`, membentuk kerangka keamanan dan pelacakan yang rapih.


## Tujuan Pembelajaran
- Middleware logging
- Auth middleware
- Context propagation

## Daftar File
- `middleware/logger.go`
- `middleware/auth.go`
- `handler/user.go`

## Cara Menjalankan
```bash
cd journey-13-middleware-context
go run main.go
```

## Catatan
Context key type safety penting.

## Referensi
- https://pkg.go.dev/context
- https://gobyexample.com/context