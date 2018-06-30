FROM golang:alpine as builder

WORKDIR /go/src/github.com/lolibrary/lolibrary/app/filesrv

COPY main.go main.go

RUN go build -ldflags="-s -w"

FROM alpine:latest as prod

COPY --from=builder /go/src/github.com/lolibrary/lolibrary/app/filesrv/filesrv /var/filesrv

WORKDIR /srv

EXPOSE 3000

CMD ["/var/filesrv"]
