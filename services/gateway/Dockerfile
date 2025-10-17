FROM golang:1.17 AS builder

ENV GO111MODULE=on

#ENV LIB_SERVICE_URL "127.0.0.1:50051"
#ENV RATING_SERVICE_URL "127.0.0.1:50053"
#ENV RESERVATION_SERVICE_URL "127.0.0.1:50052"

WORKDIR /opt/app

COPY go.mod ./
RUN go mod tidy

COPY . .

EXPOSE 8080

RUN go build -o gateway ./cmd/main.go

CMD ["./gateway"]