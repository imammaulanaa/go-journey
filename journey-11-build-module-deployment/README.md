# Journey 11: Build & Module Deployment
> **Cerita**  
Gopher menancapkan tiang bendera `go.mod`, memanggil paket dari kejauhan dengan `go get`, dan merakit aplikasi menjadi binary utuh via `go build`. Ia menulis skrip `build.sh` sebagai mantra otomatisasi, memastikan setiap build siap dikirim ke medan eksekusi.


## Tujuan Pembelajaran
- go mod init
- go get package
- go build binary
- skrip make

## Daftar File
- `build/dockerfile`
- `cmd/app/main.go`
- `configs/config.yaml`
- `internal/core/logic.go`
- `pkg/config/config.go`
- `makefile`

## Cara Menjalankan
```bash
cd journey-11-build-module-deployment
make build
```

## Catatan
Pastikan go.mod terisi.

## Referensi
- https://go.dev/doc/modules
- https://pkg.go.dev/cmd/go