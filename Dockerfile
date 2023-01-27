# syntax=docker/dockerfile:1

FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
EXPOSE 8080
