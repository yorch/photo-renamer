# Builder
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY src/go.mod src/go.sum ./
RUN go get -d -v ./...
COPY src/ .
RUN go build -o renamer main.go

# Final image
FROM alpine:latest
LABEL maintainer="jorge.barnaby@gmail.com"

COPY --from=builder /app/renamer /

ENTRYPOINT [ "/renamer" ]
