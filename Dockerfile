FROM golang:1.15 as build

MAINTAINER Myroslav Bilohubka <m.bilohubka@protonmail.com>

COPY cmd/olimp /go/src/github.com/myroslav-b/olimp/cmd/olimp
WORKDIR /go/src/github.com/myroslav-b/olimp/cmd/olimp

RUN \
    go get github.com/go-chi/chi && \
    go get github.com/go-chi/chi/middleware && \
    go get github.com/go-chi/render && \
    go get github.com/jessevdk/go-flags && \
    go get github.com/pkg/errors

RUN \
    go build -o olimp .

FROM ubuntu:20.04

RUN \
    apt-get update && \
    apt-get install ca-certificates -y

COPY --from=build /go/src/github.com/myroslav-b/olimp/cmd/olimp/olimp /srv/olimp

WORKDIR /srv

EXPOSE 8080

CMD ["/srv/olimp"]