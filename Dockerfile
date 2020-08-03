FROM golang:1.13-alpine AS builder
WORKDIR /go/src/app
COPY . . 
RUN go get -d -v ./...
RUN go build -v -o QuickMemeManager

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/app/QuickMemeManager .
CMD ["./QuickMemeManager"]
