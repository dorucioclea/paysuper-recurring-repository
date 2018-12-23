FROM golang:1.11.2-alpine AS builder

RUN apk add bash ca-certificates git

WORKDIR /application

ENV GO111MODULE=on
ENV MICRO_REGISTRY_ADDRESS=p1pay-consul
ENV MONGO_HOST=""
ENV MONGO_DB=""
ENV MONGO_USER=""
ENV MONGO_PASSWORD=""

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -o $GOPATH/bin/payone_repository .

ENTRYPOINT $GOPATH/bin/payone_repository