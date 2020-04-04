FROM golang:alpine
LABEL maintainer="jorge.barnaby@gmail.com"

RUN apk add --no-cache git

WORKDIR /app

COPY src/renamer.go .

RUN go get -d -v ./... \
    && go build renamer.go

ENTRYPOINT [ "./renamer" ]
