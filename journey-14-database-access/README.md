# Contoh SQL Table
```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  age INT
);
```

# Cara Menjalankan
```bash
go run main.go
```
coba dengan:
```bash
curl http://localhost:8080/users
```
```bash
curl -X POST http://localhost:8080/users/create \
-H "Content-Type: application/json" \
-d '{"name": "Imam", "age": 30}'
```