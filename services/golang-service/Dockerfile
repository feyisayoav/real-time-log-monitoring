FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o service main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/service .

EXPOSE 8080
CMD ["./service"]
