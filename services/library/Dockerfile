FROM golang:1.24 AS builder

ENV GO111MODULE=on


WORKDIR /opt/app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

EXPOSE 50051

RUN go build -o library ./cmd/grpc_main.go

CMD ["./library"]