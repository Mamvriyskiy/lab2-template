FROM golang

COPY ./ /src/flight
WORKDIR /flight
RUN go mod tidy && go mod vendor && go build -o ./flight/ flight/cmd/main.go

ENTRYPOINT ["flight/cmd/main"]
