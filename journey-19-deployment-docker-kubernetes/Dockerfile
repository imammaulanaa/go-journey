# Stage 1 - build
FROM golang:1.21-alpine AS builder
WORKDIR /app

COPY . .
RUN go build -o app cmd/main.go

# Stage 2 - runtime
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .

EXPOSE 8080
ENTRYPOINT ["./app"]
