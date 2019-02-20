FROM golang:1.11-alpine AS builder

RUN apk add bash ca-certificates git

WORKDIR /application

ENV GO111MODULE=on
ENV MICRO_REGISTRY=mdns
ENV MONGO_HOST=""
ENV MONGO_DB=""
ENV MONGO_USER=""
ENV MONGO_PASSWORD=""

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -o $GOPATH/bin/paysuper_repository .

ENTRYPOINT $GOPATH/bin/paysuper_repository