FROM golang:1.21-alpine3.18 as build

WORKDIR /autograding

COPY go.mod go.sum ./
COPY test ./test
COPY metrics ./metrics

RUN apk update && apk add --no-cache build-base=0.5-r3
RUN go mod download
RUN go install github.com/cucumber/godog/cmd/godog@latest
RUN cd test/godogs && go test -v -c -o functional_tests

FROM alpine:3.8
COPY --from=build /autograding/test/godogs/functional_tests /tests/functional_tests
COPY test/godogs/features /tests/features

WORKDIR /tests

CMD ["/tests/functional_tests"]
