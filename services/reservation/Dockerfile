FROM golang:1.17 AS builder

ENV GO111MODULE=on

WORKDIR /opt/app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

EXPOSE 50052

RUN go build -o reservation ./cmd/grpc_main.go

CMD ["./reservation"]