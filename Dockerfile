FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /build
COPY . .
RUN go mod tidy
RUN go build -o sgt main.go

FROM frolvlad/alpine-glibc
COPY --from=builder /build/words /usr/share/dict/words
COPY --from=builder /build/sgt .