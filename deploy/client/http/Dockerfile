FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/main /app/services/nameNode/cmd/http/main.go

EXPOSE 443

CMD ["/app/main"]