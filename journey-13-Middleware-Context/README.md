# Cara Menjalankan
```bash
go run main.go
```
Lalu coba:
- Tanpa auth header
```bash
curl http://localhost:8080/users
→ 401 Unauthorized
```
- Dengan auth header
```bash
curl -H "X-Auth-Token: valid-token" http://localhost:8080/users
→ Hello, Imam! ...
```