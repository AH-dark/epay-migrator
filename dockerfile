FROM golang:1.21.0-alpine AS builder
WORKDIR /go/src/github.com/AH-dark/epay-migrator

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o epay-migrator .

# Path: dockerfile
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app/

COPY --from=builder /go/src/github.com/AH-dark/epay-migrator/epay-migrator /app/

ENTRYPOINT ["/app/epay-migrator"]
