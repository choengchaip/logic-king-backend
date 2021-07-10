FROM golang:1.16.5-apline3.13

RUN apk update && apk upgrade && \
apk add --no-cache bash git
RUN apk add build-base
RUN go get -u github.com/pilu/fresh

WORKDIR /app
