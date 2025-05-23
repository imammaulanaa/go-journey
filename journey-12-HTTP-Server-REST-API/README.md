# Cara Menjalankan
```bash
go run main.go
```
Lalu coba:
```bash
GET http://localhost:8080/users


POST http://localhost:8080/users/create
```
Body JSON:

```json
{
  "id": 1,
  "name": "Imam",
  "age": 30
}
```