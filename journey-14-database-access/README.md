# Journey 14: Database Access
> **Cerita**  
Di tambang PostgreSQL, Gopher menggali record, menulis transaksi, dan memanggil prosedur lewat `database/sql`. Ia menyiapkan prepared statement untuk keamanan, merencanakan rollback bila runtuh, dan merayakan setiap baris data yang berhasil diambil.


## Tujuan Pembelajaran
- Connect DB
- Query & Scan
- Prepared stmt
- Transaction

## Daftar File
- `config/db.go`
- `model/user.go`
- `handler/user.go`
- `main.go`

## Cara Menjalankan
```bash
cd journey-14-database-access
go run main.go
```

## Catatan
Pastikan DB berjalan dan DSN benar.

## Referensi
- https://pkg.go.dev/database/sql
- https://github.com/lib/pq