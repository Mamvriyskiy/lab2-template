FROM golang

COPY ./ /src/bonus
WORKDIR /bonus
RUN go mod tidy && go mod vendor && go build -o ./bonus/ bonus/cmd/main.go

ENTRYPOINT ["bonus/cmd/main"]
