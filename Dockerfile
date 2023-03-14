## build
FROM golang:1.18-buster AS build-env

ADD . /go/src/devopsci

WORKDIR /go/src/devopsci

RUN make build

## run
FROM alpine:3.9


ADD conf /devopsci/conf

RUN mkdir -p /devopsci && mkdir -p /startdevopsci/log

WORKDIR /devopsci

COPY --from=build-env /go/src/devopsci/devopsci /devopsci

ENV PATH $PATH:/devopsci

EXPOSE 8080
CMD ["./devopsci"]
