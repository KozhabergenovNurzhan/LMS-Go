FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/api/main.go


FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/server .

# godotenv.Load() expects a .env file to exist; env vars are injected by docker-compose
RUN touch .env

EXPOSE 8080

CMD ["./server"]
