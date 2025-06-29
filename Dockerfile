# Builder
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY src/go.mod src/go.sum ./
RUN go mod download
COPY src/ .
RUN go build -o renamer main.go

# Final image
FROM alpine:latest
LABEL maintainer="jorge.barnaby@gmail.com"

COPY --from=builder /app/renamer /

ENTRYPOINT [ "/renamer" ]
