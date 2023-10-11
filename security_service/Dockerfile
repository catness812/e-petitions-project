FROM golang:1.21 as builder

WORKDIR /app

COPY internal/ ./internal/
COPY pkg/ ./pkg/
COPY .env ./
COPY config.yml ./
COPY go.mod ./
COPY go.sum ./
COPY main.go ./

RUN go mod download && \
    go build -o main

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/main ./

EXPOSE 9002

CMD ["./main"]