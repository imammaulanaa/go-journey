# Journey 10: CLI App & Tooling
> **Cerita**  
Di kota CLI, Gopher menempa pedang `flag`, merangkai argumen dasar dan boolean hingga perintah sederhana menari di terminal. Tak berhenti di situ, ia memasuki kuil Cobra, menyusun subcommand dan help message, menciptakan alat baris perintah yang tangguh dan modular.


## Tujuan Pembelajaran
- Flag dasar & boolean flags
- Setup Cobra CLI

## Daftar File
- `flag_basic.go`
- `flag_bool.go`
- `cobra-cli/main.go`
- `cobra-cli/cmd/root.go`

## Cara Menjalankan
```bash
cd journey-10-cli-app-tooling
go run flag_basic.go -nama=Gopher
```

## Catatan
go get github.com/spf13/cobra@latest

## Referensi
- https://pkg.go.dev/flag
- https://github.com/spf13/cobra