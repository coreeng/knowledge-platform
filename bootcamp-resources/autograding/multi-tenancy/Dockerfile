FROM golang:1.21-alpine3.18

RUN apk update && apk add --no-cache build-base

WORKDIR /autograding

RUN go install github.com/cucumber/godog/cmd/godog@latest

COPY go.mod go.sum ./

RUN go mod download

COPY test ./test

RUN go test -c ./test/acceptance/.

CMD ["sh", "-c", "./acceptance.test"]
