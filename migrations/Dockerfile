ARG DOCKER_PROXY_HOST_1=docker.io
FROM $DOCKER_PROXY_HOST_1/golang:1.20.1 as builder
ARG GIT_TAG
ARG BUILD_TIME
ARG GIT_USER
ARG GIT_TOKEN

ENV CGO_ENABLED="0"
ENV GO111MODULE="on"


RUN touch ~/.netrc
RUN echo "machine gitlab.tn.ru login $GIT_USER password $GIT_TOKEN" > ~/.netrc

WORKDIR /app
COPY ../../../migrations/migration .
RUN go version
RUN go mod download
RUN go mod verify
RUN GOOS=linux go build -tags '-trimpath' -ldflags "-s -w -extldflags '-static' -X main.gitversion=$GIT_TAG -X main.gitbuildtime=$BUILD_TIME" -o sign-migration cmd/sign-migration/*.go

FROM $DOCKER_PROXY_HOST_1/ubuntu:22.04

ENV container="docker"
ENV DEBIAN_FRONTEND="noninteractive"
ENV TERM="xterm"
ENV LANG="en_US.UTF-8"
ENV TZ="Europe/Moscow"

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN sed -i "s/archive.ubuntu.com/mirror.yandex.ru/" /etc/apt/sources.list
RUN apt update && apt upgrade -y && apt install --no-install-recommends -y

WORKDIR /app
COPY --from=builder /app/sign-migration .

CMD ["./sign-migration"]
