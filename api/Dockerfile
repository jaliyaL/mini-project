# Stage 1: Build Go binary
FROM golang:1.22 AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o server main.go

# Stage 2: Run minimal image
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]
