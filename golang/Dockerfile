FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum main.go ./

RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o server main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

RUN chmod +x /app/server

EXPOSE 1234
CMD ["/app/server"]