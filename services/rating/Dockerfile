FROM golang:1.24 AS builder

ENV GO111MODULE=on

WORKDIR /opt/app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

EXPOSE 50053

RUN go build -o rating ./cmd/grpc_main.go

CMD ["./rating"]