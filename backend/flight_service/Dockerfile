FROM golang

COPY ./ /backend

ENV DB_HOST=dpg-d3g3kdu3jp1c73flehqg-a.oregon-postgres.render.com
ENV DB_PORT=5432
ENV DB_USER=mypostgres
ENV DB_PASSWORD=cQ3yLsL0eZ6zIb6jiwOptzqDIMnb1vUT
ENV DB_NAME=mypostgres_3lod

WORKDIR /backend
RUN go mod tidy && go mod vendor && go build -o ./flight_service/ flight_service/main.go

ENTRYPOINT ["/backend/flight_service/main"]