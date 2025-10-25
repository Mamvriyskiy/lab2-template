FROM golang

COPY ./ /src/gateway
WORKDIR /gateway
RUN go mod tidy && go mod vendor && go build -o ./gateway/ gateway/cmd/main.go

ENTRYPOINT ["gateway/cmd/main"]
