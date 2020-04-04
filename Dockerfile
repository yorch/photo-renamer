# Builder
FROM golang:alpine as builder
RUN apk add --no-cache git
WORKDIR /app
COPY src/ .
RUN go get -d -v ./... \
    && go build -o renamer main.go

# Final image
FROM alpine:latest
LABEL maintainer="jorge.barnaby@gmail.com"

COPY --from=builder /app/renamer /

ENTRYPOINT [ "/renamer" ]
