FROM golang as build

COPY . .

RUN go test ./currency/... && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
        go build -a -tags netgo -ldflags '-w' \
            -o /usr/bin/local/currency ./cmd/currency/main.go

FROM scratch

COPY /usr/bin/local/currency /bin/currency

CMD /bin/currency
