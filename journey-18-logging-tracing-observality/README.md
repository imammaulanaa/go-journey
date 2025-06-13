# Journey 18: Logging, Tracing & Observability
> **Cerita**  
Gopher menyalakan lentera Logrus, menangkap jejak timestamp, level, dan fields penting. Ia menyusuri aliran trace via OpenTelemetry, mengirim span ke Jaeger, dan menumpuk metrik di Prometheus agar kerajaan Go selalu terpantau.


## Tujuan Pembelajaran
- Structured logging
- OpenTelemetry tracing
- Prometheus metrics

## Daftar File
- `logutil/log.go`
- `metrics/prometheus.go`
- `tracing/jaeger.go`
- `main.go`

## Cara Menjalankan
```bash
cd journey-18-logging-tracing-observality
go run main.go
```

## Catatan
Jalankan Jaeger & Prometheus via Docker.

## Referensi
- https://otel.io
- https://prometheus.io