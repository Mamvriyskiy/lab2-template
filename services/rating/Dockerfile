FROM golang:1.17 AS builder

ENV GO111MODULE=on

WORKDIR /opt/app

COPY . .

EXPOSE 50053

RUN go mod tidy

RUN go build -o rating ./cmd/grpc_main.go

CMD ["./rating"]