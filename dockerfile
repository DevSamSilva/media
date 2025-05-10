# Etapa 1: Build
FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go build -o media main.go

# Etapa 2: Runtime
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/media .

CMD ["./media"]
