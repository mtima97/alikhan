# Stage 1: Build
FROM golang:1.24.1-alpine as builder

ENV GO111MODULE=on GOPROXY=https://proxy.golang.org,direct

WORKDIR /app

COPY . .

RUN go build -ldflags="-s -w" -o bin/app ./cmd/app

# Stage 2: Final
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/app /app/app

EXPOSE 8080

CMD ["./app"]
