FROM golang:1.19-alpine3.16

EXPOSE 3000

WORKDIR /go/src/

RUN go run main.go