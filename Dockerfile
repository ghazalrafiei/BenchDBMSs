FROM golang:1.13

WORKDIR /opt/

RUN go clean -modcache

ENV GO111MODULE=on

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .