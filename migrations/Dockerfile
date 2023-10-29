# build stage
FROM golang:1.21 as builder
ENV GO111MODULE=on

WORKDIR /migrations

COPY ./ ./

RUN go version
RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./main .

FROM ubuntu:22.04

WORKDIR /

COPY --from=builder ./migrations/main main

ENTRYPOINT ["./main"]
