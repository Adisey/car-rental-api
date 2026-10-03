FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o app ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]