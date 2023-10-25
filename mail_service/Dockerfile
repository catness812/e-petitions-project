FROM golang:1.21.1-alpine3.17 as builder

WORKDIR /app
COPY internal ./internal
COPY pkg ./pkg
COPY templates ./templates
COPY config.yml go.mod go.sum main.go .env ./
RUN go mod download && \
    go build -o main

FROM alpine:3.17
WORKDIR /app/mail_service
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/main /app/config.yml /app/.env ./


CMD ["./main"]


#docker build -t e-petitions-mail:1.0 .