FROM golang:1.19-alpine3.17

RUN apk update && apk add --no-cache build-base

WORKDIR /acceptance-criteria

COPY go.mod go.sum ./

RUN go mod download

COPY test ./test
COPY metrics ./metrics

RUN go install github.com/cucumber/godog/cmd/godog@latest

CMD ["sh", "-c", "cd test/godogs && go test -v"]
