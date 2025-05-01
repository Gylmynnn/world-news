FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/main.go

FROM debian:bookworm-slim

RUN useradd -m appuser

WORKDIR /app

COPY --from=builder /app/server .

USER appuser

EXPOSE 3000

CMD ["./server"]
