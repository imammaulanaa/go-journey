

# Journey 10: CLI App & Tooling

Belajar membangun aplikasi CLI dengan `flag` dan `cobra`.

## 🎯 Tujuan Pembelajaran
- Membuat CLI menggunakan package `flag`
- Menambahkan flag boolean dan string
- Membuat aplikasi CLI modular menggunakan `cobra`

## 📦 Daftar File

| File                        | Topik                        |
|-----------------------------|------------------------------|
| flag_basic.go            | CLI dengan flag              |
| flag_bool.go             | Boolean flag                 |
| cobra-cli/main.go           | Aplikasi CLI menggunakan Cobra |
| cobra-cli/cmd/root.go       | Root command Cobra           |

## ▶️ Cara Menjalankan

```bash
go run flag_basic.go -nama=Imam -umur=30
```
Untuk cobra:
```bash
go run cobra-cli/main.go
```