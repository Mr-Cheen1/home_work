FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o server ./server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

CMD ["/app/server"]