FROM golang:alpine

WORKDIR /src

COPY . .

RUN go install ./...
