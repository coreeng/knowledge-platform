FROM golang:1.21-alpine3.18 as builder

RUN apk update && apk add --no-cache build-base && apk add curl

WORKDIR /autograding

COPY go.mod go.sum ./

RUN go mod download

COPY test ./test
RUN go test -c ./test/acceptance/.

FROM alpine:3
COPY --from=builder /autograding/acceptance.test .

CMD ["sh", "-c", "./acceptance.test"]