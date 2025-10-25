FROM golang

COPY ./ /src/ticket
WORKDIR /ticket
RUN go mod tidy && go mod vendor && go build -o ./ticket/ ticket/cmd/main.go

ENTRYPOINT ["ticket/cmd/main"]
