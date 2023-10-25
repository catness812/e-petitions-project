FROM golang:1.21-alpine3.17 as builder

WORKDIR /app

COPY internal/ ./internal/
COPY pkg/ ./pkg/
COPY .env config.yml go.mod go.sum main.go ./

RUN go mod download && \
    go build -o main

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app/main /app/config.yml /app/.env ./

EXPOSE 50051

CMD ["./main"]