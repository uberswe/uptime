FROM golang:latest

ENV URLS="https://ma.rkus.io,https://www.uberswe.com" \
    INTERVAL=60

WORKDIR /usr/src/app

COPY ./cmd/uptime/ .

ENTRYPOINT go run ./main.go -interval=${INTERVAL} -urls=${URLS}