FROM golang:1.21-alpine3.18 as builder

RUN apk add --no-cache build-base=0.5-r3 && apk add --no-cache curl=8.4.0-r0

WORKDIR /autograding

COPY go.mod go.sum ./

RUN go mod download

COPY test ./test
RUN go test -c ./test/acceptance/.

FROM alpine:3

WORKDIR /
COPY --from=builder /autograding/acceptance.test .

CMD ["sh", "-c", "./acceptance.test"]