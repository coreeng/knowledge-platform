FROM golang:1.21-alpine3.18

RUN apk update && apk add --no-cache build-base && apk add curl

WORKDIR /autograding

RUN go install github.com/cucumber/godog/cmd/godog@latest

COPY go.mod go.sum ./

RUN go mod download

RUN curl -LO "https://dl.k8s.io/release/v1.26.1/bin/linux/amd64/kubectl"
RUN install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

COPY test ./test
RUN go test -c ./test/acceptance/.
CMD ["sh", "-c", "./acceptance.test"]
