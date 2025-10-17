FROM golang:1.17 AS builder

ENV GO111MODULE=on

ENV DATABASE_URL "postgres://ujctenexixakdj:e77e07688d8f4489a7a5e2324616405018b6d1387fdaa9176dabd81370582848@ec2-79-125-30-28.eu-west-1.compute.amazonaws.com:5432/dbpu8qgisgfd9f"

WORKDIR /opt/app

COPY . .

EXPOSE 50053

RUN go mod tidy

RUN go build -o rating ./cmd/grpc_main.go

CMD ["./rating"]