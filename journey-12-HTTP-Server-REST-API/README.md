# Journey 12: HTTP Server & REST API
> **Cerita**  
Di gerbang HTTP, Gopher membuka jalur GET dan POST, menyusun handler yang merespons permintaan dengan JSON murni. Dengan `http.HandleFunc`, ia menata rute, memanggil model data, dan mengemas respons lewat `encoding/json`, menjamin setiap pesan tiba dengan sempurna.


## 🎯 Tujuan Pembelajaran
- HTTP server
- JSON request/response
- Error handling

## 📂 Daftar File
- `main.go`
- `handler/user.go`
- `model/user.go`
- `utils/response.go`

## ▶️ Cara Menjalankan
```bash
cd journey-12-http-server-rest-api
go run main.go
```

## ⚠️ Catatan
Gunakan Content-Type application/json.

## 📚 Referensi
- https://pkg.go.dev/net/http
- https://pkg.go.dev/encoding/json