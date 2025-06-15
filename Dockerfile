FROM golang:1.19

WORKDIR /app

RUN go mod init calculator

COPY . .

RUN go build -o calculator ./cmd/main.go

CMD ["./calculator"]