FROM golang:alpine as build

ENV GOPATH /go

WORKDIR /go/src/github.com/lolibrary/lolibrary/api

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep && dep ensure -v

RUN go generate ./... && go test ./...

RUN go build -o /var/api main.go
