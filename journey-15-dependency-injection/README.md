# Journey 15: Dependency Injection
> **Cerita**  
Gopher menelusuri benang dependensi, menyulam lapisan service dan repository dengan constructor injection. Ia membangun interface abstraction, memudahkan switch implementasi, dan membuka jalan bagi mocking di medan pengujian.


## Tujuan Pembelajaran
- Constructor injection
- Interface abstraction
- Decoupling

## Daftar File
- `repository/user_repo.go`
- `service/user_service.go`
- `model/user.go`
- `main.go`

## Cara Menjalankan
```bash
cd journey-15-dependency-injection
go run main.go
```

## Catatan
Mock repo untuk testing.

## Referensi
- https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis