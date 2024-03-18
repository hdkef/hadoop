FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/main /app/services/nameNode/cmd/grpc/main.go

FROM golang:1.20

WORKDIR /app

COPY --from=builder /app/main /app/main

EXPOSE 443

CMD ["/app/main"]