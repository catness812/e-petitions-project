FROM golang:1.21.3-alpine3.18 as builder
WORKDIR /app

COPY internal ./internal
COPY pkg ./pkg
COPY templates ./templates
COPY config.yml go.mod go.sum main.go .env ./

RUN go mod download && \
    go build -o main

FROM alpine:3.18
WORKDIR /app/mail_service

RUN apk add openssl && \
    adduser -D user && chown -R user /app

COPY --from=builder /app/templates ./templates
COPY --from=builder /app/main /app/config.yml /app/.env ./

USER user

CMD ["./main"]