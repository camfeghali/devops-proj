FROM golang:alpine

RUN apk update && apk --no-cache add curl && apk add --no-cache bash
