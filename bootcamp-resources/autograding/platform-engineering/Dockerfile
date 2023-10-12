FROM golang:1.21-alpine3.18 as builder

RUN apk update && apk add --no-cache build-base && apk add curl

WORKDIR /autograding

COPY go.mod go.sum ./

RUN go mod download

COPY test ./test
RUN go test -c ./test/acceptance/.

FROM alpine:3

COPY --from=builder /autograding/acceptance.test .

#The manifests are used by kubectl directly in the tests, they need to be copied as well
COPY test/acceptance/manifests /test/acceptance/manifests

RUN apk add curl
RUN curl -LO "https://dl.k8s.io/release/v1.26.1/bin/linux/amd64/kubectl"
RUN install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

CMD ["sh", "-c", "./acceptance.test"]
