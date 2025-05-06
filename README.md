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
- **Benchmarking**: Mengukur performa aplikasi dengan Go.
- **Error Handling**: Mengelola error dengan baik, termasuk menggunakan custom error types dan error wrapping.

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

### **journey 8: Concurrency & Parallelism**
- **Goroutine**: Memahami cara kerja goroutine untuk menjalankan fungsi secara bersamaan.
- **Channel**: Menggunakan channel untuk komunikasi antar goroutine.
- **Mutex & RWMutex**: Mengontrol akses data bersama dengan menggunakan Mutex.
- **Select Statement**: Mengelola beberapa channel dalam concurrency.

### **journey 9: Deployment dan Containerization**
- **Docker**: Menggunakan Docker untuk mengemas aplikasi Go ke dalam container.
- **CI/CD Pipeline**: Integrasi aplikasi Go ke dalam pipeline CI/CD menggunakan GitHub Actions atau GitLab CI.
- **Kubernetes**: Men-deploy aplikasi Go dalam cluster Kubernetes.

### **journey 10: Microservices dengan Golang**
- **gRPC**: Menggunakan gRPC untuk komunikasi antar layanan microservices.
- **Protocol Buffers**: Menggunakan Protocol Buffers untuk serialisasi data.
- **Service Discovery & Load Balancing**: Memahami cara kerja layanan dan balancing di arsitektur microservices.
- **API Gateway**: Mengelola layanan-layanan mikro dengan API Gateway.

### **journey 11: Pembuatan Command-Line Tools**
- **Cobra**: Membuat dan mengelola aplikasi command-line menggunakan Cobra.
- **Flag Parsing**: Menggunakan `flag` package untuk parsing argumen di CLI.
- **File System**: Menangani file system melalui CLI tools.

### **journey 12: Membuat Library/Package untuk Digunakan Kembali**
- **Membuat Package**: Membangun library Go untuk digunakan dalam proyek lain.
- **GoDoc**: Dokumentasi kode dengan GoDoc.
- **Distribusi Library**: Mengelola distribusi library Go di GitHub atau repository lainnya.

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
├── journey-8-concurrency/
├── journey-9-deployment/
├── journey-10-microservices/
├── journey-11-cli-tools/
├── journey-12-library-package/
```
---

## 🧰 Tools & Referensi
- [Go by Example](https://gobyexample.com/)
- [Tour of Go](https://tour.golang.org/)
- [Go.dev](https://go.dev/)
- [Go Playground](https://play.golang.org/)
