FROM golang:1.17

RUN apt-get update && apt -y install curl libpq-dev
RUN go get -u github.com/cosmtrek/air

WORKDIR /app

EXPOSE 8080
