FROM golang:1.21-alpine3.18

RUN apk update && apk add --no-cache build-base=0.5-r3

WORKDIR /autograding

COPY go.mod go.sum ./

RUN go mod download

COPY test ./test
COPY metrics ./metrics
COPY structs ./structs
COPY consts ./consts

RUN go install github.com/cucumber/godog/cmd/godog@latest

CMD ["sh", "-c", "cd test/godogs && go test -v"]
