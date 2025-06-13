# Journey 16: Internal Package & Modules
> **Cerita**  
Gopher menutup perpustakaan rahasia di `internal/`, menyembunyikan kode privat dari dunia luar. Di sisi lain, ia menyusun modul reusable di `pkg/`, menyebarluaskan utility untuk rekan seperjalanan di seluruh kerajaan Go.


## Tujuan Pembelajaran
- internal/ usage
- pkg/ for public modules
- modular structure

## Daftar File
- `internal/service/math_service.go`
- `internal/repository/user_repo.go`
- `pkg/utils/string_utils.go`
- `cmd/app/main.go`

## Cara Menjalankan
```bash
cd journey-16-internal-package-modules
go run cmd/app/main.go
```

## Catatan
Kode di internal tidak bisa import luar modul.

## Referensi
- https://go.dev/doc/go1.4#internalpackages