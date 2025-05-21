FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go build -o app ./cmd/app

FROM alpine:latest
COPY --from=builder /app/app /app
ENTRYPOINT ["/app"]
