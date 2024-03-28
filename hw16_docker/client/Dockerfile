FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./  
RUN go mod download

COPY client/ ./client/

WORKDIR /app/client

RUN go build -o /app/client ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/client /app/

CMD ["./client"]