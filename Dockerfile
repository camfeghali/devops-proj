FROM golang:alpine

RUN apk update && apk --no-cache add curl && apk add --no-cache bash

RUN mkdir app

ADD app /app

WORKDIR /app

RUN go build -o main .

# CMD ["/app/main"]

