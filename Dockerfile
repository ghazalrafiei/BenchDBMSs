FROM golang:1.13

WORKDIR $GOPATH/src/github.com/ghazalrafiei/BenchDBMSs

RUN go clean -modcache

ENV GO111MODULE=on

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .
EXPOSE 8080

CMD ["./main"]