# 🚀 Go Journey

Repository ini berisi materi dan contoh kode untuk mempelajari bahasa pemrograman Go (Golang), mulai dari dasar hingga topik lanjutan. Setiap journey mencakup contoh kode yang dapat digunakan untuk memahami konsep-konsep utama dalam Golang.

## 🧭 Roadmap Pembelajaran

### **journey 1: Pengenalan Dasar**
- **Hello World**: Memahami cara kerja program pertama di Go.
- **Variabel & Tipe Data**: Belajar mendeklarasikan variabel dan berbagai tipe data di Go.
- **Operator**: Menggunakan operator aritmatika, logika, dan perbandingan.
- **Kontrol Eksekusi**: Penggunaan statement `if`, `switch`, dan `for`.
- **Fungsi**: Membuat fungsi sederhana dengan return value.

### **journey 2: Struktur Data & Kontrol Lanjut**
- **Array, Slice, Map**: Memahami tipe data koleksi dan cara penggunaannya.
- **Struct & Method**: Belajar mendeklarasikan dan menggunakan struct, serta menambahkan method pada struct.
- **Pointer**: Penggunaan pointer untuk manajemen memori yang efisien.

### **journey 3: Fitur Khusus Go**
- **Interface**: Memahami konsep interface dan penggunaannya dalam Go.
- **Goroutine & Channel**: Mengelola concurrency dengan goroutine dan komunikasi antar goroutine menggunakan channel.

### **journey 4: Testing & Best Practices**
- **Unit Testing**: Menulis dan menjalankan unit test dengan `testing` package.

### **journey 5: Golang Standard Library**
- **`net/http`**: Membuat HTTP server dan client.
- **`io/ioutil` & `os`**: Bekerja dengan file system dan I/O.
- **`strings` & `strconv`**: Manipulasi string dan konversi tipe data.
- **`time`**: Mengelola waktu dan tanggal.
- **`sync`**: Sinkronisasi goroutine menggunakan mutex.
- **`encoding/json`**: Mengolah data JSON.
- **`log`**: Untuk mencatat log aplikasi.

### **journey 6: Web Development dengan Go**
- **RESTful API**: Membangun RESTful API dengan Go menggunakan `net/http`.
- **Routing**: Menggunakan routing dengan **Gorilla Mux** atau **Echo**.
- **Middleware**: Menambahkan autentikasi dan logging dalam aplikasi web.
- **HTML Templates**: Menggunakan template untuk render HTML pada server-side.

### **journey 7: Error Handling dan Custom Errors**
- **Custom Error Types**: Membuat dan menangani error khusus dengan struct.
- **Error Wrapping**: Teknik error wrapping menggunakan `fmt.Errorf`.
- **Best Practices**: Menangani error secara eksplisit dan bersih untuk aplikasi yang lebih stabil.

### **journey 8: Testing & benchmarking**
- **Benchmarking**: Mengukur performa aplikasi dengan Go.
- **Mocking**: Mocking interface untuk test
- **table test**: Table-driven test
- **unit test**: Unit test

### **journey 9: Concurrency & Parallelism**
- **Goroutine**: Memahami cara kerja goroutine untuk menjalankan fungsi secara bersamaan.
- **Channel**: Menggunakan channel untuk komunikasi antar goroutine.
- **Mutex & RWMutex**: Mengontrol akses data bersama dengan menggunakan Mutex.
- **Select Statement**: Mengelola beberapa channel dalam concurrency.

### **journey 10: CLI APP & tTooling**
- **flag_basic**: CLI dasar dengan flag
- **flag_bool**: Boolean flag
- **cobra-cli**: CLI modern menggunakan Cobra

### **journey 11: Build Module Deployment**
- **mod init**: Memahami inisialisasi dan penggunaan modul dalam Go.
- **test package**: Menggunakan third-party package dalam proyek Go.
- **build binary**: Memahami cara membangun aplikasi Go menjadi binary executable.
- **app-structure/main.go**: Memahami struktur aplikasi nyata dalam proyek Go.
- **build.sh**: Menggunakan skrip build sederhana untuk membangun dan mendistribusikan aplikasi.

### **journey 12: HTTP Server & REST API dengan net/http**
- **HTTP Server**: Membuat server HTTP sederhana untuk menangani permintaan.
- **RESTful API**: Membangun RESTful API yang mengembalikan data dalam format JSON.
- **Query Parameters**: Menangani query parameters dalam permintaan HTTP.
- **HTTP Methods**: Menangani berbagai metode HTTP (GET, POST, PUT, DELETE).
- **Error Handling**: Menangani error dalam server HTTP.

### **journey 13: Middleware & Context**
- **Middleware**: Menambahkan middleware untuk autentikasi dan logging dalam aplikasi web.
- **Context**: Mengelola request dengan context untuk memberikan batas waktu dan informasi tambahan.
- **Middleware untuk Autentikasi**: Menambahkan middleware untuk autentikasi pengguna sebelum mengakses endpoint tertentu.
- **Context untuk Mengelola Data**: Menggunakan context untuk mengelola data yang dapat dibagikan antar fungsi dalam satu permintaan.

### **journey 14: Database Access**
- **Database Connection**: Mengakses database menggunakan Go dengan package database/sql.
- **Inserting Data**: Menyisipkan data ke dalam database menggunakan perintah SQL.
- **Updating Data**: Memperbarui data yang ada dalam database.
- **Deleting Data**: Menghapus data dari database.
- **Error Handling in Database Operations**: Menangani error saat berinteraksi dengan database.

### **journey 15: Dependency Injection**
- **Dependency Injection**: Menerapkan dependency injection di Go untuk meningkatkan modularitas dan pengujian.

### **journey 16: Deployment ke Docker / Kubernetes**
- **Docker**: Membuat Dockerfile untuk aplikasi Go.
- **Kubernetes**: Mendefinisikan deployment Kubernetes untuk aplikasi Go.

### **journey 17: Membuat Internal Package & Module**
- **Internal Package**: Membuat internal package di Go untuk menjaga encapsulation.

### **journey 18: Error Handling yang Lebih Advance**
- **Advanced Error Handling**: Menangani error dengan lebih baik menggunakan custom error.
- **Error Type**: Membuat dan menggunakan custom error types.

### **journey 19: Logging, Tracing, dan Observability**
- **Logging**: Mencatat log aplikasi untuk keperluan debugging dan monitoring.
- **Tracing**: Melakukan tracing untuk aplikasi untuk memahami alur eksekusi.
### **journey 20: Microservice Architecture (Intro)**

---

## 📂 Struktur Folder
```hcl
go-journey/
├── journey-1-dasar/
├── journey-2-struktur-data/
├── journey-3-fitur-go/
├── journey-4-testing-best-practice/
├── journey-5-standard-library/
├── journey-6-web-development/
├── journey-7-error-handling/
├── journey-8-testing-benchmarking/
├── journey-9-concurency-parallesim/
├── journey-10-cli-app-tools/
├── journey-11-build-module-deployment/
├── journey-12-http-server-rest-api/
├── journey-13-middleware-context/
├── journey-14-database-access/
├── journey-15-dependency-injection/
├── journey-16-deployment-docker-kubernetes/
├── journey-17-internal-package-modules/
├── journey-18-error-handling-advanced/
├── journey-19-logging-tracing-observality/
├── journey-20-microservice-architecture/
```
---

## 🧰 Tools & Referensi
- [Go by Example](https://gobyexample.com/)
- [Tour of Go](https://tour.golang.org/)
- [Go.dev](https://go.dev/)
- [Go Playground](https://play.golang.org/)
