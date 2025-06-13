# Journey 8: Testing & Benchmarking
> **Cerita**  
Di laboratorium pengujian, Gopher merancang eksperimen table-driven test layaknya resep ilmiah, memproduksi mock sederhana untuk meniru dependensi, dan mengukur performa fungsi dengan benchmark. Ia menyadari, hanya lewat pengujian mendalam, kerajaan Go-nya bisa berdiri kuat.


## Tujuan Pembelajaran
- Table-driven tests
- Benchmark tests
- Mocking interface

## Daftar File
- `unit_test.go`
- `table_test.go`
- `benchmark.go`
- `mocking.go`

## Cara Menjalankan
```bash
cd journey-8-testing-benchmarking
go test -bench=.
```

## Catatan
Mock sederhana cukup stub type.

## Referensi
- https://pkg.go.dev/testing
- https://github.com/stretchr/testify