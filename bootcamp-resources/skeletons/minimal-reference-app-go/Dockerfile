FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY . /app

RUN go build -o server /app/cmd/main.go

FROM alpine:3.18.4

COPY --from=builder /app/server ./

CMD ["./server"]