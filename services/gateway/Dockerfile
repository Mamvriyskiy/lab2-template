FROM golang:1.24 AS builder

ENV GO111MODULE=on

#ENV LIB_SERVICE_URL "127.0.0.1:8060"
#ENV RATING_SERVICE_URL "127.0.0.1:8050"
#ENV RESERVATION_SERVICE_URL "127.0.0.1:8070"

WORKDIR /opt/app

COPY go.mod ./
RUN go mod tidy

COPY . .

EXPOSE 8080

RUN go build -o gateway ./cmd/main.go

CMD ["./gateway"]