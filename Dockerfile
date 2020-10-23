FROM golang:1.15-alpine AS builder
LABEL org.opencontainers.image.source=https://github.com/chand1012/QuickMemeManager
WORKDIR /go/src/app
COPY . . 
RUN go get -d -v ./...
RUN go build -v -o QuickMemeManager

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/app/QuickMemeManager .
CMD ["./QuickMemeManager"]
